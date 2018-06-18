import React from 'react'

export default class Actions extends React.Component {
    render() {
        return (
            <div className="actions">
                {this.props.children}
            </div>
        )
    }
}
