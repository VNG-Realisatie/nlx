// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

/**
 * Log key...val pairs as log group
 * @param {Object} data: key - value pairs. title is used to display title of group
 * @param {String} data.title: required title (string) *
 */
export const logGroup = (data) => {
    let props = Object.keys(data)

    let ignore = []

    if (data.title) {
        // eslint-disable-next-line
        console.group(data.title);
        ignore.push('title')
    } else {
        // eslint-disable-next-line
        console.group("logGroup");
    }

    props.forEach((key) => {
        // only if not in ignore list
        if (ignore.indexOf(key) === -1) {
            // eslint-disable-next-line
            console.log(key, "...", data[key]);
        }
    })
    // eslint-disable-next-line
    console.groupEnd();
}

export default logGroup
