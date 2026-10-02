import type { ChatMessage } from '../../types/chat'

import MessageList from './MessageList'
import MessageInput from './MessageInput'

interface ChatWindowProps {
  messages: ChatMessage[]
  message: string
  error: string
  isSending: boolean
  onMessageChange: (message: string) => void
  onSendMessage: () => void
}

function ChatWindow({
  messages,
  message,
  error,
  isSending,
  onMessageChange,
  onSendMessage,
}: ChatWindowProps) {
  return (
    <main className="chat-area">
      <div className="chat-content">
        {messages.length === 0 ? (
          <div className="welcome-message">
            <h1>How can I help?</h1>
            <p>Ask VeerAI anything.</p>
          </div>
        ) : (
          <MessageList messages={messages} />
        )}
      </div>

      {error !== '' && (
        <p className="chat-error">{error}</p>
      )}

      <MessageInput
        message={message}
        isSending={isSending}
        onMessageChange={onMessageChange}
        onSendMessage={onSendMessage}
      />
    </main>
  )
}

export default ChatWindow