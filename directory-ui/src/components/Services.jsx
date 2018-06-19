import React from 'react'
import Service from './Service'

export default class Services extends React.Component {
    render() {
        const {
            serviceList,
            sortBy,
            sortAscending
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

        const sortArrow = (
            <svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
                <g id="arrow-down" fill="none" fillRule="evenodd">
                    <path id="Shape" fill="currentColor" fillRule="nonzero" transform={sortAscending ? "rotate(90 4 5)" : "rotate(-90 4 5)"} d="M5 4h-6v2h6v3l4-4-4-4z"></path>
                </g>
            </svg>
        )

        return (
            <div className="table-responsive">
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            <th scope="col" className={sortBy === 'inway_addresses' ? "sorting" : ""}>
                                <button onClick={(e) => this.props.onSort('inway_addresses')}>
                                    Status
                                    {sortBy === 'inway_addresses' && sortArrow}
                                </button>
                            </th>
                            <th scope="col" className={sortBy === 'organization_name' ? "sorting" : ""}>
                                <button onClick={(e) => this.props.onSort('organization_name')}>
                                    Organisation
                                    {sortBy === 'organization_name' && sortArrow}
                                </button>
                            </th>
                            <th scope="col" className={sortBy === 'name' ? "sorting" : ""}>
                                <button onClick={(e) => this.props.onSort('name')}>
                                    Service
                                    {sortBy === 'name' && sortArrow}
                                </button>
                            </th>
                            <th scope="col">
                                <button disabled>
                                    API address
                                </button>
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
