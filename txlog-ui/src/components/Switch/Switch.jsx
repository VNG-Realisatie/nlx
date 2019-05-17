// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import classnames from 'classnames'

class Switch extends Component {
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

export default Switch
