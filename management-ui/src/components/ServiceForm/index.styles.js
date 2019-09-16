// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import styled from 'styled-components'
import theme from '../../theme'

export const StyledFormGroupColumnContainer = styled.div`
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;

    @media screen and (min-width: ${theme.breakpoints.breakpoints.sm}px) {
        margin: 0 -28px;
    }
`

export const StyledFormGroupColumn = styled.div`
    flex: 1 0 100%;
    padding: 0;

    @media screen and (min-width: ${theme.breakpoints.breakpoints.sm}px) {
        padding: 0 28px;
        flex: 1 1 50%;
    }
`

export const StyledFormGroup = styled.div`
    margin-bottom: 1rem;
`

export const StyledButtonGroup = styled.div`
    display: flex;
    justify-content: space-between;
`

export const StyledDeletableField = styled.div`
    display: flex;
    margin-bottom: 5px;

    button {
        margin-left: 10px;
    }
`
