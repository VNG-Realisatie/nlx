import React, { Component } from 'react'

class Mock extends Component {
    state = {
        counter: 0,
    }
    render() {
        return (
            <div data-test-id="mock-component">
                <h1>This is mock component for testing</h1>
            </div>
        )
    }
}

export default Mock
