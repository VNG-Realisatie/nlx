// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { shallow } from 'enzyme'
import ServicesTableContainer from './ServicesTableContainer'

describe('ServicesTableContainer', () => {
    let wrapper
    let instance

    beforeEach(() => {
        wrapper = shallow(<ServicesTableContainer />)
        instance = wrapper.instance()
    })

    describe('sorting', () => {
        describe('when no sort properties are provided', () => {
            it('should have null values for default sorting', () => {
                expect(wrapper.state('sortBy')).toBeNull()
                expect(wrapper.state('sortOrder')).toBeNull()
            })
        })

        describe('when sort properties are provided', () => {
            it('should have the sorting state set to those values', () => {
                wrapper = shallow(
                    <ServicesTableContainer
                        sortBy="test-column"
                        sortOrder="test-order"
                    />,
                )
                expect(wrapper.state('sortBy')).toEqual('test-column')
                expect(wrapper.state('sortOrder')).toEqual('test-order')
            })
        })

        describe('sorting the services', () => {
            it('should sort by property in descending order', () => {
                const services = [
                    { name: 'abc' },
                    { name: 'def' },
                    { name: 'ghi' },
                ]

                const result = instance.sortServices(services, 'name', 'desc')
                expect(result).toEqual([
                    { name: 'ghi' },
                    { name: 'def' },
                    { name: 'abc' },
                ])
            })

            it('should be case insensitive', () => {
                const services = [
                    { name: 'def' },
                    { name: 'GHI' },
                    { name: 'ABC' },
                ]

                const result = instance.sortServices(services, 'name')
                expect(result).toEqual([
                    { name: 'ABC' },
                    { name: 'def' },
                    { name: 'GHI' },
                ])
            })
        })

        describe('toggle the sorting', () => {
            describe('when no sorting was active before', () => {
                it('should sort ascending', () => {
                    instance.onToggleSorting('organization')
                    expect(wrapper.state('sortBy')).toBe('organization')
                    expect(wrapper.state('sortOrder')).toBe('asc')
                })
            })

            describe('when the column was already sorted on', () => {
                it('should reverse the sorting', () => {
                    instance.onToggleSorting('organization')
                    expect(wrapper.state('sortOrder')).toBe('asc')

                    instance.onToggleSorting('organization')
                    expect(wrapper.state('sortOrder')).toBe('desc')

                    instance.onToggleSorting('organization')
                    expect(wrapper.state('sortOrder')).toBe('asc')
                })
            })
        })
    })

    describe('filtering', () => {
        describe('by status', () => {
            it('should filter out the offline services', () => {
                const services = [{ status: 'online' }, { status: 'offline' }]

                const result = instance.filterServicesByOnlineStatus(services)
                expect(result).toHaveLength(1)
            })
        })

        describe('by a query', () => {
            it('should filter on organization name', () => {
                const services = [
                    { organization: 'abc', name: 'def' },
                    { organization: 'ghi', name: 'jkl' },
                ]

                const result = instance.filterServicesByQuery(services, 'abc')
                expect(result).toHaveLength(1)
            })

            it('should filter on service name', () => {
                const services = [
                    { organization: 'abc', name: 'def' },
                    { organization: 'ghi', name: 'jkl' },
                ]

                const result = instance.filterServicesByQuery(services, 'def')
                expect(result).toHaveLength(1)
            })
        })
    })
})
