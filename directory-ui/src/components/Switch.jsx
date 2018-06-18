import React from 'react'

export default class Search extends React.Component {
    render() {
        const {
            id,
            children
        } = this.props

        return (
            <div className="form-switch">
                <input type="checkbox" className="form-switch-input" id={id} />
                <label className="form-switch-label" htmlFor={id}>{children}</label>
            </div>
        )
    }
}
