// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'

const SortArrow = ({ sortAscending }) => (
    <svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
        <g id="arrow-down" fill="none" fillRule="evenodd">
            <path
                id="Shape"
                fill="currentColor"
                fillRule="nonzero"
                transform={sortAscending ? 'rotate(90 4 5)' : 'rotate(-90 4 5)'}
                d="M5 4h-6v2h6v3l4-4-4-4z"
            />
        </g>
    </svg>
)

export default SortArrow
