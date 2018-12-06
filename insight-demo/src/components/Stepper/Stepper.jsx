import React, { PureComponent } from 'react'
import { NavLink } from 'react-router-dom';
import styled from 'styled-components'

const Wrapper = styled.div`
    position: relative;
    display: inline-flex;
`

const gutterWidth = 96
const progressWidth = gutterWidth + 32

class Stepper extends PureComponent {
    getProgressWidth = (pathname) => {
        switch (pathname) {
            case '/':
                return 0
            case '/stepone':
                return 0
            case '/steptwo':
                return progressWidth
            case '/stepthree':
                return progressWidth * 2
            case '/stepfour':
                return progressWidth * 3
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
                width: ${p => `${this.getProgressWidth(this.props.pathname)}px`};
                background-color: ${p => p.theme.color.primary.main};
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
            font-size: 16px;
            font-weight: bold;
            text-decoration: none;

            user-select: none;
            box-sizing: border-box;

            &:not(:last-child) {
                margin-right: ${`${gutterWidth}px`};
            }

            &[aria-current] {
                color: ${p => p.theme.color.primary.main};
                border-color: ${p => p.theme.color.primary.main};
            }
        `

        return (
            <Wrapper>
                <Progress/>
                <Step as={NavLink} to="/stepone" done={['/steptwo', '/stepthree', '/stepfour'].includes(this.props.pathname)}>1</Step>
                <Step as={NavLink} to="/steptwo" done={['/stepthree', '/stepfour'].includes(this.props.pathname)}>2</Step>
                <Step as={NavLink} to="/stepthree" done={['/stepfour'].includes(this.props.pathname)}>3</Step>
                <Step as={NavLink} to="/stepfour">4</Step>
            </Wrapper>
        )
    }
}

Stepper.propTypes = {
}

Stepper.defaultProps = {
}

export default Stepper
