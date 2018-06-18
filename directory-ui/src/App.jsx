import React, { Component } from 'react'
import Navigation from './components/Navigation'
import Actions from './components/Actions'
import Search from './components/Search'
import Switch from './components/Switch'
import Services from './components/Services'

class App extends Component {
    constructor(props) {
        super(props)

        this.state = {
            q: ''
        }

        this.onChange = this.onChange.bind(this)
    }

    onChange(e) {
        this.setState({ q: e.target.value })
    }

    render() {
        return (
            <div className="App">
                <Navigation />
                <Actions>
                    <div className="container">
                        <div className="row">
                            <div className="col-sm-6 col-lg-4 offset-lg-2">
                                <Search onChange={this.onChange} value={this.state.q} />
                            </div>
                            <div className="col-sm-6 col-lg-6 d-flex align-items-center">
                                <Switch id="switch1">Switch me</Switch>
                            </div>
                        </div>
                    </div>
                </Actions>
                <section>
                    <div className="container">
                        <Services q={this.state.q} />
                    </div>
                </section>
            </div>
        );
    }
}

export default App;
