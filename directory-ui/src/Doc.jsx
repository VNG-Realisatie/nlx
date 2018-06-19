import React from 'react'
import { RedocStandalone } from 'redoc'
import './static/css/redoc-override.css';

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
                            family: '"Muli", sans-serif',
                            smoothing: 'antialiased',
                            optimizeSpeed: true,
                        },
                        headingsFont: {
                            family: '"Muli", sans-serif',
                        },
                        code: {
                            fontSize: '14px',
                            fontFamily: '"Fira Code", monospaced',
                        },
                        colors: {
                            main: '#3d83fa',
                            success: '#00aa13',
                            redirect: '#ffa500',
                            error: '#e53935',
                            info: '#87ceeb',
                            text: '#263238',
                            code: '#e83e8c',
                            codeBg: '#f8f9fa',
                            warning: '#f1c400',
                            http: {
                                get: '#6bbd5b',
                                post: '#248fb2',
                                put: '#9b708b',
                                options: '#d3ca12',
                                patch: '#e09d43',
                                delete: '#e27a7a',
                                basic: '#999',
                                link: '#31bbb6',
                                head: '#c167e4',
                            }
                        },
                        menu: {
                            width: '260px',
                            backgroundColor: '#f8f9fa',
                        },
                        rightPanel: {
                            backgroundColor: '#ffffff',
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
