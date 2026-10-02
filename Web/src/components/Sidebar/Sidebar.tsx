import { useEffect, useState } from 'react'

import { getConversations } from '../../services/api'

interface Conversation {
  id: number
  title: string
  created_at: string
}

interface SidebarProps {
  onNewChat: () => void
  onSelectConversation: (conversationId: number) => void
  onDeleteConversation: (conversationId: number) => Promise<void>
  activeConversationId: number | null
  conversationRefreshKey: number
  
}

function Sidebar({
  onNewChat,
  onSelectConversation,
  onDeleteConversation,
  activeConversationId,
  conversationRefreshKey,
}: SidebarProps) {
  const [conversations, setConversations] = useState<Conversation[]>([])

  useEffect(() => {
    loadConversations()
  }, [conversationRefreshKey])

  async function loadConversations() {
    const token = localStorage.getItem('veerai_token')

    if (!token) {
      return
    }

    try {
      const existingConversations = await getConversations(token)

      setConversations(existingConversations)
    } catch {
      console.error('Could not load conversations.')
    }
  }

  return (
    <aside className="sidebar">
      <div className="sidebar-header">
        <h2>VeerAI</h2>
      </div>

      <button
        className="new-chat-button"
        onClick={onNewChat}
      >
        <span>+</span>
        New Chat
      </button>

      <div className="conversation-list">
        {conversations.map((conversation) => (
          <div
            className={`conversation-item ${conversation.id === activeConversationId ? 'active' : ''
              }`}
            key={conversation.id}
          >
            <button
              className="conversation-select"
              onClick={() => onSelectConversation(conversation.id)}
            >
            {conversation.title}
            </button>

            <button
              className="conversation-delete"
              onClick={async () => {
                const confirmed = window.confirm(
                  'Delete this conversation?'
                )

                if (!confirmed) {
                  return
                }

                await onDeleteConversation(conversation.id)

                setConversations((currentConversations) =>
                  currentConversations.filter(
                    (item) => item.id !== conversation.id
                  )
                )
              }}
            >
              🗑️
            </button>
          </div>
        ))}
      </div>
    </aside>
  )
}

export default Sidebar