// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import styled from 'styled-components'

export const StyledServiceDetailPane = styled.div`
  position: fixed;
  right: 0;
  top: 56px;
  width: 300px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45, 50, 64, 0.05), 0 1px 8px rgba(45, 50, 64, 0.05);
  z-index: 1;
  min-height: calc(100% - 56px);
  padding: 20px 24px;
`

export const StyledHeader = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
`

export const StyledTitle = styled.h3`
  color: #517fff;
  font-size: 20px;
  line-height: 28px;
  font-weight: 700;
  margin-bottom: 0;
`

export const StyledSecondTitle = styled.h4`
  display: block;
  color: #a3aabf;
  font-size: 16px;
  line-height: 28px;
  font-weight: 500;
  margin: 0;
`

export const StyledCloseButton = styled.button`
  width: 40px;
  height: 40px;
  background: none;
  border: none;
  border-radius: 50%;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;

  &:hover {
    background-color: #f7f9fc;
  }

  &:active {
    background-color: #f0f2f7;
  }
`

export const StyledSubtitle = styled.h4`
  font-weight: 600;
  font-size: 12px;
  line-height: 20px;
  color: #2d3240;
  margin-bottom: 0;
`

export const StyledValue = styled.p`
  margin-left: 0;
  margin-top: 0;
  padding-bottom: 8px;
  font-size: 12px;
  line-height: 20px;
  color: #2d3240;
`

export const StyledLink = styled.a`
  color: #2d3240;
  text-decoration: underline;
`
