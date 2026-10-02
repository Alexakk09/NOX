import { useState } from 'react'

import type { ChatMessage } from '../types/chat'

import {
  createConversation,
  getConversationMessages,
  sendChatMessage,
  deleteConversation,
} from '../services/api'

function useChat(onConversationChanged?: () => void) {
  const [conversationId, setConversationId] =
    useState<number | null>(null)

  const [currentMessage, setCurrentMessage] = useState('')

  const [chatMessages, setChatMessages] =
    useState<ChatMessage[]>([])

  const [chatError, setChatError] = useState('')

  const [isSendingMessage, setIsSendingMessage] =
    useState(false)

  async function startConversation(token: string) {
    const newConversationId = await createConversation(token)

    setConversationId(newConversationId)
    onConversationChanged?.()
  }

  async function selectConversation(selectedConversationId: number) {
    const token = localStorage.getItem('veerai_token')

    if (!token) {
      setChatError('You are not logged in.')
      return
    }

    setChatError('')

    try {
      const existingMessages = await getConversationMessages(
        token,
        selectedConversationId,
      )

      setConversationId(selectedConversationId)
      setChatMessages(existingMessages)
      setCurrentMessage('')
    } catch {
      setChatError('Could not load conversation.')
    }
  }

  async function sendMessage() {
    const trimmedMessage = currentMessage.trim()

    if (trimmedMessage === '') {
      return
    }

    const token = localStorage.getItem('veerai_token')

    if (!token) {
      setChatError('You are not logged in.')
      return
    }

    if (conversationId === null) {
      setChatError('No conversation is active.')
      return
    }

    setChatError('')
    setIsSendingMessage(true)

    const userMessage: ChatMessage = {
      content: trimmedMessage,
      sender: 'user',
    }

    setChatMessages((previousMessages) => [
      ...previousMessages,
      userMessage,
    ])

    setCurrentMessage('')

    try {
      const response = await sendChatMessage(
        token,
        conversationId,
        trimmedMessage,
      )

      const assistantMessage: ChatMessage = {
        content: response,
        sender: 'assistant',
      }

      setChatMessages((previousMessages) => [
        ...previousMessages,
        assistantMessage,
      ])
      onConversationChanged?.()
    } catch {
      setChatError('Could not get a response from VeerAI.')
    } finally {
      setIsSendingMessage(false)
    }
  }

  async function removeConversation(selectedConversationId: number) {
    const token = localStorage.getItem('veerai_token')

    if (!token) {
      setChatError('You are not logged in.')
      return
    }

    try {
      await deleteConversation(token, selectedConversationId)

      if (selectedConversationId === conversationId) {
        setConversationId(null)
        setChatMessages([])
        setCurrentMessage('')
      }
    } catch {
      setChatError('Could not delete conversation.')
    }
  }

  async function startNewChat() {
    const token = localStorage.getItem('veerai_token')

    if (!token) {
      return
    }

    try {
      await startConversation(token)

      setChatMessages([])
      setCurrentMessage('')
      setChatError('')
    } catch {
      setChatError('Could not create a new conversation.')
    }
  }

  return {
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
    removeConversation
  }
}

export default useChat