import ReactMarkdown from 'react-markdown'
import type { ChatMessage } from '../../types/chat'

interface MessageListProps {
  messages: ChatMessage[]
}

function MessageList({ messages }: MessageListProps) {
  return (
    <div className="message-list">
      {messages.map((chatMessage, messageIndex) => (
        <div
          className={`message ${chatMessage.sender}`}
          key={messageIndex}
        >
          <ReactMarkdown>
            {chatMessage.content}
          </ReactMarkdown>
        </div>
      ))}
    </div>
  )
}

export default MessageList