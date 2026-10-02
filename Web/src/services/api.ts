import type {
  ChatMessage,
  ConversationResponse,
  ChatResponse,
} from '../types/chat'

const API_BASE_URL = 'http://localhost:8080'

export async function login(
  username: string,
  password: string,
): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      username,
      password,
    }),
  })

  if (!response.ok) {
    throw new Error('Login failed.')
  }

  const loginData = await response.json()

  return loginData.token
}

export async function createConversation(
  token: string,
): Promise<number> {
  const response = await fetch(
    `${API_BASE_URL}/api/conversations`,
    {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
    },
  )

  if (!response.ok) {
    throw new Error('Failed to create conversation.')
  }

  const conversationData: ConversationResponse =
    await response.json()

  return conversationData.id
}

export interface Conversation {
  id: number
  title: string
  created_at: string
}

export async function getConversations(
  token: string,
): Promise<Conversation[]> {
  const response = await fetch(
    `${API_BASE_URL}/api/conversations`,
    {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    },
  )

  if (!response.ok) {
    throw new Error('Failed to get conversations.')
  }

  const conversations: Conversation[] =
    await response.json()

  return conversations
}

export async function getConversationMessages(
  token: string,
  conversationId: number,
): Promise<ChatMessage[]> {
  const response = await fetch(
    `${API_BASE_URL}/api/conversations/${conversationId}/messages`,
    {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    },
  )

  if (!response.ok) {
    throw new Error('Failed to get conversation messages.')
  }

  const messages = await response.json()

  return messages.map((message: { role: string; content: string }) => ({
    content: message.content,
    sender: message.role === 'user' ? 'user' : 'assistant',
  }))
}

export async function sendChatMessage(
  token: string,
  conversationId: number,
  message: string,
): Promise<string> {
  const response = await fetch(
    `${API_BASE_URL}/api/chat`,
    {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        conversation_id: conversationId,
        message,
      }),
    },
  )

  if (!response.ok) {
    throw new Error('Chat request failed.')
  }

  const chatData: ChatResponse = await response.json()

  return chatData.response
}

export async function deleteConversation(
  token: string,
  conversationId: number
): Promise<void> {
  const response = await fetch(
    `${API_BASE_URL}/api/conversations/${conversationId}`,
    {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    }
  )

  if (!response.ok) {
    throw new Error('Failed to delete conversation')
  }
}