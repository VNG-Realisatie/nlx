import React, { Component } from 'react'

import './Spinner.scss'

class Spinner extends Component {
    render() {
        return (
          <div className="Spinner">
              <div className="Spinner__BulletContainer">
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
                  <div className="Spinner__Bullet" />
              </div>
          </div>
        )
    }
}

export default Spinner
