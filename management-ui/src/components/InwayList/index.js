// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { Fragment } from 'react'
import { shape, arrayOf, string } from 'prop-types'
import { Card } from './index.styles'

const InwayList = ({ result }) => (
    <Fragment>
        {result.inways && result.inways.length > 0 ? (
            result.inways.map((inway) => (
                <Card key={inway.name}>{inway.name}</Card>
            ))
        ) : (
            <p>There are no inways.</p>
        )}
    </Fragment>
)

InwayList.propTypes = {
    result: shape({
        inways: arrayOf(shape({ name: string.isRequired })),
    }),
}

export default InwayList
