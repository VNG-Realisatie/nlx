// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import PropTypes from 'prop-types'
import styled from 'styled-components'

const Section = styled.div`
    padding: 3rem 0 2rem;
    background-color: ${(p) => p.backgroundColor && p.backgroundColor};
    color: ${(p) => p.color && p.color};
`

Section.propTypes = {
    backgroundColor: PropTypes.string,
    color: PropTypes.string,
}

export default Section
