// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React, { useState, useEffect } from 'react'
import { string, func, node, shape } from 'prop-types'
import UserRepository from '../domain/user-repository'

const UserContext = React.createContext()

const UserContextProvider = ({
  children,
  fetchAuthenticatedUser,
  user: defaultUser,
}) => {
  const [user, setUser] = useState(defaultUser || null)
  let canceled = false

  useEffect(
    () => {
      const fetchUser = async () => {
        try {
          const authenticatedUser = await fetchAuthenticatedUser()

          if (!canceled) {
            setUser(authenticatedUser)
          }
        } catch (err) {
          setUser(null)
        }
      }

      if (!defaultUser) {
        fetchUser()
      }
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  )

  const cancelFetchHandler = () => {
    canceled = true
  }

  const onLoginHandler = (user) => {
    setUser(user)
  }

  const onLogoutHandler = () => {
    setUser(null)
  }

  return (
    <UserContext.Provider
      value={{
        user: user,
        cancelFetch: cancelFetchHandler,
        onLoginHandler: onLoginHandler,
        onLogoutHandler: onLogoutHandler,
      }}
    >
      {children}
    </UserContext.Provider>
  )
}

UserContextProvider.propTypes = {
  fetchAuthenticatedUser: func,
  children: node,
  user: shape({
    id: string,
    fullName: string,
    email: string,
    pictureUrl: string,
  }),
}

UserContextProvider.defaultProps = {
  fetchAuthenticatedUser: UserRepository.getAuthenticatedUser,
}

export default UserContext
export { UserContextProvider }
