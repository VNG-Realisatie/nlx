module.exports = (opts = {}) => {
    return (code, state) => {
      const output = `
        import React from 'react'

        export default class ${state.componentName} extends React.Component {
            render() {
                const props = this.props

                return (
                    ${code}
                )
            }
        }
      `

      return output
    }
}
