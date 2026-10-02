export interface ChatMessage {
  content: string
  sender: 'user' | 'assistant'
}

export interface ConversationResponse {
  id: number
}

export interface ChatResponse {
  response: string
}