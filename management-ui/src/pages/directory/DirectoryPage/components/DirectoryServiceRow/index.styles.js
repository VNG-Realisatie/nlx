// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled, { css } from 'styled-components'
import Table from '../../../../../components/Table'

export const StyledTd = styled(Table.Td)`
  ${(props) =>
    props.color &&
    css`
      color: ${props.color};
    `}
`

export const StyledTdAccess = styled(Table.Td)`
  width: 18rem;
  max-width: 18rem;
  padding-top: 0;
  padding-bottom: 0;
`

export const AccessMessageWrapper = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-end;
`
