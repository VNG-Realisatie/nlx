import React from 'react'
import classnames from 'classnames'

export default class Search extends React.Component {
    render() {
        const { id, children, alwaysOn } = this.props

        return (
            <div
                className={classnames({
                    'form-switch': true,
                    'form-switch--alwaysOn': alwaysOn,
                })}
            >
                <input
                    type="checkbox"
                    className="form-switch-input"
                    id={id}
                    onChange={this.props.onChange}
                    checked={this.props.checked}
                />
                <label className="form-switch-label" htmlFor={id}>
                    {children}
                </label>
            </div>
        )
    }
}
