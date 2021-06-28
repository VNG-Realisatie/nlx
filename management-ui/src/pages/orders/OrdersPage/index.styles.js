// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'
import { Button } from '@commonground/design-system'

export const Centered = styled.section`
  display: flex;
  flex-direction: column;
  align-items: center;
  margin: 10rem auto;
  max-width: 32rem;
  text-align: center;
`

export const ActionsBar = styled.div`
  display: flex;

  > *:nth-child(3) {
    margin-left: auto;
  }
  > *:last-child {
    margin-left: 16px;
  }
`

export const StyledButton = styled(Button)`
  ${({ isActive, theme }) =>
    isActive &&
    `
    background-color: ${theme.colorBackgroundButtonSecondarySelected};
    color: ${theme.colorTextButtonSecondarySelected};

    :after {
      top: -2px;
    }
    :hover, :focus {
      background-color: ${theme.colorBackgroundButtonSecondarySelectedHover};
      color: ${theme.colorTextButtonSecondarySelected};
    }
  `}
`
