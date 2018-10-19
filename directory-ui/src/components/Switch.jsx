import React from 'react'

import './Switch-addon.css';

export default class Search extends React.Component {
    render() {
        const {
            id,
            children
        } = this.props

        return (
            <div className="form-switch">
                <input type="checkbox" className="form-switch-input" id={id} onChange={this.props.onChange} checked={this.props.checked} />
                <label className="form-switch-label" htmlFor={id}>{children}</label>
            </div>
        )
    }
}
