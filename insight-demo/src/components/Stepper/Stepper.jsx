// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { PureComponent } from 'react'
import PropTypes from 'prop-types'
import { NavLink } from 'react-router-dom';
import styled from 'styled-components'
import { media } from '../../theme/helpers'

const Wrapper = styled.div`
    position: relative;
    display: inline-flex;
`

const gutterWidth = 80
const progressWidth = gutterWidth + 32
const smallGutterWidth = 32
const smallProgressWidth = smallGutterWidth + 32

class Stepper extends PureComponent {
    getProgressWidth = (pathname, size = 'large') => {
        const width = size === 'large' ? progressWidth : smallProgressWidth
        switch (pathname) {
            case '/':
                return 0
            case '/stepone':
                return 0
            case '/steptwo':
                return width
            case '/stepthree':
                return width * 2
            case '/stepfour':
                return width * 3
            default:
                break
        }
    }

    render() {
        const Progress = styled.div`
            position: absolute;
            height: 3px;
            top: 14px;
            left: 16px;
            right: 16px;
            background-color: #EAEAEA;

            &:before {
                content: '';
                display: block;
                height: 100%;
                background-color: ${p => p.theme.color.primary.main};

                ${media.xsDown`
                    width: ${`${this.getProgressWidth(this.props.pathname, 'small')}px`};
                `}

                ${media.xsUp`
                    width: ${`${this.getProgressWidth(this.props.pathname)}px`};
                `}
            }
        `

        const Step = styled.div`
            position: relative;

            display: flex;
            align-items: center;
            justify-content: center;

            width: 32px;
            height: 32px;
            padding-bottom: 2px;
            border-radius: 50%;
            border: 3px solid transparent;
            border-color: ${p => !p.done && p.theme.color.grey[30]};

            background-color: ${p => p.done ? p.theme.color.primary.main : p.theme.color.white};
            color: ${p => p.done ? p.theme.color.white : p.theme.color.grey[60]};

            font-family: ${p => p.theme.font.family.main};
            font-size: 16px;
            font-weight: bold;
            text-decoration: none;

            user-select: none;
            box-sizing: border-box;
            transition: background-color ${p => p.theme.transition.fast}, color ${p => p.theme.transition.fast}, border-color ${p => p.theme.transition.fast};


            &:not(:last-child) {
                ${media.xsDown`
                    margin-right: ${`${smallGutterWidth}px`};
                `}

                ${media.xsUp`
                    margin-right: ${`${gutterWidth}px`};
                `}
            }

            &[aria-current] {
                color: ${p => p.theme.color.primary.main};
                border-color: ${p => p.theme.color.primary.main};
                pointer-events: none;
            }

            &:hover {
                background-color: ${p => p.done && p.theme.color.primary.light};
                border-color: ${p => !p.done && p.theme.color.grey[40]};
            }
        `

        return (
            <Wrapper>
                <Progress/>
                <Step as={NavLink} to="/stepone" done={['/steptwo', '/stepthree', '/stepfour'].includes(this.props.pathname) ? 1 : 0}>1</Step>
                <Step as={NavLink} to="/steptwo" done={['/stepthree', '/stepfour'].includes(this.props.pathname) ? 1 : 0}>2</Step>
                <Step as={NavLink} to="/stepthree" done={['/stepfour'].includes(this.props.pathname) ? 1 : 0}>3</Step>
                <Step as={NavLink} to="/stepfour">4</Step>
            </Wrapper>
        )
    }
}

Stepper.propTypes = {
    pathname: PropTypes.string.isRequired
}

export default Stepper
