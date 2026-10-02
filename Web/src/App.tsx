import { useEffect, useState } from 'react'
import './App.css'

import {  getConversations, login } from './services/api'

import LoginForm from './components/Login/LoginForm'
import Sidebar from './components/Sidebar/Sidebar'
import ChatWindow from './components/Chat/ChatWindow'

import useChat from './hooks/useChat'

function App() {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [loginError, setLoginError] = useState('')
  const [conversationRefreshKey, setConversationRefreshKey] = useState(0)

  const {
    conversationId,
    currentMessage,
    chatMessages,
    chatError,
    isSendingMessage,
    setCurrentMessage,
    sendMessage,
    startConversation,
    startNewChat,
    selectConversation,
    removeConversation,

  }  = useChat(() => {
  setConversationRefreshKey((currentKey) => currentKey + 1)
})

  useEffect(() => {
    const token = localStorage.getItem('veerai_token')

    if (!token) {
      return
    }

    setIsLoggedIn(true)

    async function loadExistingConversations(token: string) {
      try {
        const conversations = await getConversations(token)

        if (conversations.length === 0) {
          return
        }

        const latestConversation = conversations[0]

        await selectConversation(latestConversation.id)
      } catch {
        localStorage.removeItem('veerai_token')
        setIsLoggedIn(false)
      }
    }

    loadExistingConversations(token)
  }, [])

  async function handleLogin() {
    setLoginError('')

    if (username.trim() === '' || password === '') {
      setLoginError('Please enter your username and password.')
      return
    }

    try {
      const token = await login(username, password)

      localStorage.setItem('veerai_token', token)

      setIsLoggedIn(true)

      await startConversation(token)
    } catch {
      setLoginError('Login failed. Check your username and password.')
    }
  }

  if (!isLoggedIn) {
    return (
      <LoginForm
        username={username}
        password={password}
        loginError={loginError}
        onUsernameChange={setUsername}
        onPasswordChange={setPassword}
        onLogin={handleLogin}
      />
    )
  }

  return (
    <div className="app">
      <Sidebar
        onNewChat={startNewChat}
        onSelectConversation={selectConversation}
        onDeleteConversation={removeConversation}
        activeConversationId={conversationId}
        conversationRefreshKey={conversationRefreshKey}
      />

      <ChatWindow
        messages={chatMessages}
        message={currentMessage}
        error={chatError}
        isSending={isSendingMessage}
        onMessageChange={setCurrentMessage}
        onSendMessage={sendMessage}
      />
    </div>
  )
}

export default App