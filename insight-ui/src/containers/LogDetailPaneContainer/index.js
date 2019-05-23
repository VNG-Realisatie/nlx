import React, { Component } from 'react'
import {shape, string, arrayOf, instanceOf} from 'prop-types'
import { connect } from 'react-redux'
import LogDetailPane from '../../components/LogDetailPane'

class LogDetailPaneContainer extends Component {
  constructor(props) {
    super(props)

    this.onCloseHandler = this.onCloseHandler.bind(this)
  }

  onCloseHandler() {
    const { history, parentURL } = this.props
    history.push(parentURL)
  }

  render() {
    const { log } = this.props
    return log ? <LogDetailPane closeHandler={this.onCloseHandler} {...log} /> : null
  }
}

LogDetailPaneContainer.propTypes = {
  parentURL: string.isRequired,
  match: shape({
    params: shape({
      logid: string
    })
  }),
  log: shape({
    id: string,
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    application: string,
    reason: string,
    date: instanceOf(Date)
  })
}

LogDetailPaneContainer.defaultProps = {
  match: {
    params: {}
  }
}

const mapStateToProps = ({ logs }, ownProps) => {
  return {
    log: logs.records
      .find(log => log.id === ownProps.match.params.logid)
  }
}

export default connect(
  mapStateToProps
)(LogDetailPaneContainer)

