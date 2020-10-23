// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { oneOfType, string, number, bool, arrayOf, func } from 'prop-types'

const Switch = ({ test, children }) => {
  const matchedCase = children.find((child) => {
    const { value } = child.props
    return Array.isArray(value) ? value.includes(test) : value === test
  })

  if (matchedCase) return matchedCase

  const defaultCase = children.find((child) => child.type.name === 'Default')

  if (!defaultCase) {
    console.error(
      `No case matched or default found for: <Switch test={${test}} />`,
    )
  }

  return defaultCase || null
}

const Case = ({ children }) => children()
const Default = ({ children }) => children()

Switch.propTypes = {
  test: oneOfType([string, number, bool]).isRequired,
}

Case.propTypes = {
  value: oneOfType([string, number, bool, arrayOf(string, number, bool)])
    .isRequired,
  children: func, // if not function, it will always render
}

Default.propTypes = {
  children: func, // if not function, it will always render
}

Switch.Case = Case
Switch.Default = Default

export default Switch
