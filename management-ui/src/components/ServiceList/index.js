// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { Fragment } from 'react'
import { Link } from 'react-router-dom'
import { shape, arrayOf, string } from 'prop-types'
import { Card } from './index.styles'

const ServiceList = ({ result }) => (
    <Fragment>
        {result.services && result.services.length > 0 ? (
            result.services.map((service) => (
                <Card key={service.name}>
                    <Link to={`/services/update/${service.name}`}>
                        {service.name}
                    </Link>
                </Card>
            ))
        ) : (
            <p>There are no services.</p>
        )}
    </Fragment>
)

ServiceList.propTypes = {
    result: shape({
        services: arrayOf(shape({ name: string.isRequired })),
    }),
}

export default ServiceList
