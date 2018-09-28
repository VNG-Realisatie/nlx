import React, { Component } from 'react'
import { QRCode } from 'react-qr-svg'
import axios from 'axios'

let interval

export default class Irma extends Component {
    constructor(props) {
        super(props)

        this.state = {
            qrCode: "",
            error: false
        }
    }

    componentDidMount() {
        const { organization } = this.props

        axios({
            method: 'post',
            url: `http://${organization.insight_log_endpoint}:30080/generateJWT`,
            data: {
                dataSubjects: [
                    'over18'
                ]
            }
        }).then(response => {
            const firstJWT = response.data

            axios({
                method: 'post',
                url: `http://${organization.insight_irma_endpoint}:30080/api/v2/verification/`,
                headers: { 'content-type': 'text/plain' },
                data: firstJWT
            }).then(response => {
                let irmaVerificationRequest = response.data
                const u = irmaVerificationRequest['u']

                // prepend IrmaVerifiationRequest with URL
                irmaVerificationRequest['u'] = `http://${organization.insight_irma_endpoint}:30080/api/v2/verification/${u}`

                const qrCode = JSON.stringify(irmaVerificationRequest)
                console.log(qrCode)

                this.setState({ qrCode })

                interval = setInterval(() => {
                    axios({
                        method: 'get',
                        url: `http://${organization.insight_irma_endpoint}:30080/api/v2/verification/${u}/status`
                    }).then(response => {
                        //console.log(response.data)
                        this.props.afterLogin(organization, response.data)
                    }).catch(error => {
                        clearInterval(interval)
                        this.setState({ error: true })
                    })
                }, 5000)
            })
        })
    }

    componentWillUnmount() {
        clearInterval(interval)
    }

    render() {
        return (
            <div>
                Please login with Irma<br /><br />
                { this.state.error && (
                    <b>Something went wrong, refresh and try again</b>
                )}
                { this.state.qrCode &&
                    <QRCode
                        bgColor="#FFFFFF"
                        fgColor="#000000"
                        level="Q"
                        style={{ width: 256 }}
                        value={this.state.qrCode}
                    />
                }
            </div>
        )
    }
}