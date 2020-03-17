// code based on https://marmelab.com/blog/2018/07/18/accessibility-performance-testing-puppeteer.html
const axe = require('axe-core')
const { printReceived } = require('jest-matcher-utils')
const { resolve } = require('path')
const { realpathSync, mkdirSync, existsSync } = require('fs')

const PATH_TO_AXE = './node_modules/axe-core/axe.min.js';
const appDirectory = realpathSync(process.cwd());

const resolvePath = relativePath => resolve(appDirectory, relativePath);

exports.analyzeAccessibility = async (page, screenshotPath, options = {}) => {
    // Inject the axe script in our page
    await page.addScriptTag({ path: resolvePath(PATH_TO_AXE) });
    // we make sure that axe is executed in the next tick after
    // the page emits the load event, giving priority for the
    // original JS to be evaluated
    const accessibilityReport = await page.evaluate(axeOptions => {
        return new Promise(resolve => {
            setTimeout(resolve, 0);
        }).then(() => axe.run(axeOptions));
    }, options);

    if (
        screenshotPath &&
        (accessibilityReport.violations.length ||
            accessibilityReport.incomplete.length)
    ) {
        const path = `${process.cwd()}/screenshots`;
        if (!existsSync(path)) {
            mkdirSync(path);
        }
        await page.screenshot({
            path: `${path}/${screenshotPath}`,
            fullPage: true,
        });
    }

    return accessibilityReport;
};

const defaultOptions = {
    violationsThreshold: 0,
    incompleteThreshold: 0,
};

const printInvalidNode = node =>
    `- ${printReceived(node.html)}\n\t${
        node.any.map(check => check.message).join('\n\t')
        }`;

const printInvalidRule = rule =>
    `${printReceived(rule.help)} on ${
        rule.nodes.length
        } nodes\r\n${rule.nodes
        .map(printInvalidNode)
        .join('\n')}`;


// Add a new method to expect assertions with a very detailed error report
expect.extend({
    toHaveNoAccessibilityIssues(accessibilityReport, options) {
        let violations = [];
        let incomplete = [];
        const finalOptions = Object.assign({}, defaultOptions, options);

        if (
            accessibilityReport.violations.length >
            finalOptions.violationsThreshold
        ) {
            violations = [
                `Expected to have no more than ${
                    finalOptions.violationsThreshold
                    } violations. Detected ${
                    accessibilityReport.violations.length
                    } violations:\n`,
            ].concat(accessibilityReport.violations.map(printInvalidRule));
        }

        if (
            finalOptions.incompleteThreshold !== false &&
            accessibilityReport.incomplete.length >
            finalOptions.incompleteThreshold
        ) {
            incomplete = [
                `Expected to have no more than ${
                    finalOptions.incompleteThreshold
                    } incomplete. Detected ${
                    accessibilityReport.incomplete.length
                    } incomplete:\n`,
            ].concat(accessibilityReport.incomplete.map(printInvalidRule));
        }

        const message = [].concat(violations, incomplete).join('\n');
        const pass =
            accessibilityReport.violations.length <=
            finalOptions.violationsThreshold &&
            (finalOptions.incompleteThreshold === false ||
                accessibilityReport.incomplete.length <=
                finalOptions.incompleteThreshold);

        return {
            pass,
            message: () => message,
        };
    },
});
