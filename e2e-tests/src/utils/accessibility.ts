/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

import AxeBuilder from "@axe-core/webdriverjs";
import webdriver from "selenium-webdriver";

interface Violation {
  id: string;
  impact: string;
  description: string;
  nodes: Array<any>;
}

export const checkForAccessibilityIssues = async (
  driver: webdriver.ThenableWebDriver,
  rulesToDisable: Array<string>
): Promise<Array<any>> => {
  const builder = await new AxeBuilder(driver);
  const results = await builder.disableRules(rulesToDisable).analyze();
  return logViolations(results.violations);
};

export const logViolations = (violations: Array<any>) => {
  return violations.map(({ id, impact, description, nodes }: Violation) => ({
    id,
    impact,
    description,
    nodes: nodes.length,
  }));
};
