import React from 'react'
import Service from './Service'

export default class Services extends React.Component {
    render() {
        const {
            serviceList,
            sortBy,
            sortAscending
        } = this.props

        const services = serviceList.map((data) => (
            <Service
                key={data.organization_name + data.service_name}
                data={data}
            />
        ))

        const sortArrow = (
            <svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
                <g id="arrow-down" fill="none" fillRule="evenodd">
                    <path id="Shape" fill="currentColor" fillRule="nonzero" transform={sortAscending ? "rotate(90 4 5)" : "rotate(-90 4 5)"} d="M5 4h-6v2h6v3l4-4-4-4z"></path>
                </g>
            </svg>
        )

        const centerStyle = {
            textAlign: 'center'
        }

        return (
            <div className="table-responsive mb-5">
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            <th scope="col" className={sortBy === 'inway_addresses' ? "sorting" : ""}>
                                <button style={centerStyle} onClick={(e) => this.props.onSort('inway_addresses')}>
                                    Status
                                    {sortBy === 'inway_addresses' && sortArrow}
                                </button>
                            </th>
                            <th scope="col" className={sortBy === 'organization_name' ? "sorting" : ""}>
                                <button onClick={(e) => this.props.onSort('organization_name')}>
                                    Organization
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
                                    API Type
                                </button>
                            </th>
                            <th scope="col">
                                <button style={centerStyle} disabled>
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
