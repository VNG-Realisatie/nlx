// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import styled from 'styled-components'

export const Wrapper = styled.div`
  position: relative;
  display: flex;
`

export const Input = styled.input`
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  margin: 0;
  opacity: 0;

  &[disabled] + label {
    color: #b4b4b4;
  }

  &:not([disabled]) {
    cursor: pointer;

    &:checked + label {
      &:before {
        background-color: #30709d;
      }
      &:after {
        border-color: #30709d;
        transform: translateX(14px);
      }
    }
  }
`

export const Label = styled.label`
  padding: 0 0 0 ${(p) => p.theme.tokens.spacing09};
  line-height: 22px;
  user-select: none;
  pointer-events: none;

  &:before {
    content: '';
    position: absolute;
    left: 0;
    width: ${(p) => p.theme.tokens.spacing08};
    height: ${(p) => p.theme.tokens.spacing06};
    border: 0;
    border-radius: 12px;
    background-color: ${({ theme }) => theme.tokens.colorPaletteGray300};
    transition: background-color 0.25s ease;
  }

  &:after {
    content: '';
    position: absolute;
    width: 20px;
    height: 20px;
    top: 2px;
    left: 4px;
    border-radius: 50%;
    background-color: white;
    box-shadow: -2px 1px 1px 0px rgba(0, 0, 0, 0.15);
    transform: translateX(-1px);
    transition: transform 0.25s ease;
  }
`
