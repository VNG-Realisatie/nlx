import React from 'react'
import { RedocStandalone } from 'redoc'
import './style.css';

export default class Doc extends React.Component {
    render() {
        return (
            <RedocStandalone
                specUrl={"https://api.apis.guru/v2/specs/instagram.com/1.0.0/swagger.yaml"}
                options={{
                    theme: {
                        baseFont: {
                            size: '14px',
                            lineHeight: '1.5',
                            weight: '300',
                            family: 'Muli, sans-serif',
                            smoothing: 'antialiased',
                            optimizeSpeed: true,
                        },
                        headingsFont: {
                            family: 'Muli, sans-serif',
                        },
                        code: {
                            fontSize: '13px',
                            fontFamily: 'Fira Code, monospace',
                        },
                        colors: {
                            main: '#3d83fa'
                        },
                        menu: {
                            width: '260px',
                            backgroundColor: '#f8f9fa',
                        },
                        rightPanel: {
                            backgroundColor: '#263238',
                            width: '40%',
                        }
                    },
                    scrollYOffset: '56',
                    hideLoading: true
                }}
            />
        )
    }
}
