// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
export default function cssUnitOrEmpty(props, propName) {
  // eslint-disable-next-line security/detect-object-injection
  if (props[propName] === undefined) return

  // eslint-disable-next-line security/detect-object-injection
  if (!/(^|px|rem|%)$/.test(props[propName])) {
    return new Error(
      `Invalid prop \`${propName}\` supplied to \`Modal\`. Use a css value in px, rem or %.`,
    )
  }
}
