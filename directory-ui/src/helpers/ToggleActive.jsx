import React from 'react';
import tippy from 'tippy.js';
import './TooltipTippy.scss';

export default class TooltipTippy extends React.Component {
    componentDidMount() {
        const firstChild = this.refs.firstChild.getWrappedInstance ? this.refs.firstChild.getWrappedInstance() : this.refs.firstChild

        this.tippy = tippy(firstChild, {
            placement: this.props.placement || 'bottom',
            duration: [500,500],
            html: false,
            trigger: 'click'
        })
    }

    componentWillUnmount() {
        if (this.tippy) {
            this.tippy.destroyAll()
        }
    }

    render() {
        if (Array.isArray(this.props.children)) {
            console.error("You are trying to add a tooltip to multiple children, please use tooltip on one child only.")
            return this.props.children
        }

        return React.cloneElement(this.props.children, {
            ref: 'firstChild',
            title: this.props.title
        })
    }
}