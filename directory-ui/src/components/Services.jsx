import React from 'react'
import Service from './Service'
import axios from 'axios';

export default class Services extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            services: []
        }
    }

    componentDidMount() {
        axios.get(`/api/directory/list-services`)
            .then(res => {
                const services = res.data.services;
                this.setState({ services });
            })
            .catch(e => {
                this.errors.push(e)
            })
    }

    render() {
        const {
            serviceFilter
        } = this.props

        const filteredServices = this.state.services.filter((service) => {
            if (!serviceFilter) {
                return true
            }

            return service.name.toLowerCase().includes(serviceFilter.toLowerCase()) ||
                service.organization_name.toLowerCase().includes(serviceFilter.toLowerCase())
        })

        const services = filteredServices.map((service) => (
            <Service
                key={service.organization_name + service.name}
                organizationName={service.organization_name}
                name={service.name}
                inwayAddresses={service.inway_addresses}
                documentationUrl={service.documentation_url}
            />
        ))

        return (
            <div className="table-responsive">
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            <th scope="col" className="sorting ascending">
                                <button>
                                    Status
                                    <svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
                                        <g id="arrow-down" fill="none" fillRule="evenodd">
                                            <path id="Shape" fill="currentColor" fillRule="nonzero" transform="rotate(90 4 5)" d="M5 4h-6v2h6v3l4-4-4-4z"></path>
                                        </g>
                                    </svg>
                                </button>
                            </th>
                            <th scope="col">
                                <button>Organisation</button>
                            </th>
                            <th scope="col">
                                <button>Service</button>
                            </th>
                            <th scope="col">
                                <button>API</button>
                            </th>
                        </tr>
                    </thead>
                    <tbody>
                        {services}
                    </tbody>
                </table>
            </div>
        )
    }
}
