import React from 'react'
import Service from './Service'

export default class Services extends React.Component {
    render() {
        const {
            serviceList
        } = this.props

        const services = serviceList.map((service) => (
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
