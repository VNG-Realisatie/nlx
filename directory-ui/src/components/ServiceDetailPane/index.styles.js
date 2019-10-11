// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import styled from 'styled-components'

export const StyledServiceDetailPane = styled.div`
  position: fixed;
  right: 0;
  top: 56px;
  width: 300px;
  background: #ffffff;
  box-shadow: 0 0 0 1px rgba(45,50,64,.05), 0 1px 8px rgba(45,50,64,.05);
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
  color: #517FFF;
  font-size: 20px;
  line-height: 28px;
  font-weight: 700;
  margin-bottom: 0;
`

export const StyledSecondTitle = styled.h4`
  display: block;
  color: #A3AABF;
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
    background-color: #F7F9FC;
  }

  &:active {
    background-color: #F0F2F7;
  }
`

export const StyledSubtitle = styled.h4`
  font-weight: 600;
  font-size: 12px;
  line-height: 20px;
  color: #2D3240;
  margin-bottom: 8px;
`

export const StyledDl = styled.dl`
  margin-top: 0;
  font-size: 12px;
  line-height: 20px;
  overflow: hidden;

  dt {
    color: #A3AABF;
    width: 90px;
    float: left;
    clear: both;
    padding-bottom: 8px;
  }

  dd {
    color: #2D3240;
    float: right;
    margin-left: 0;
    width: calc(100% - 90px);
    padding-bottom: 8px;
  }
`

export const StyledEmailAddressLink = styled.a`
  color: #2D3240;
  text-decoration: underline;
`