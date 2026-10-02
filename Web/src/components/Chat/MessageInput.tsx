interface MessageInputProps {
  message: string
  isSending: boolean
  onMessageChange: (message: string) => void
  onSendMessage: () => void
}

function MessageInput({
  message,
  isSending,
  onMessageChange,
  onSendMessage,
}: MessageInputProps) {
  return (
    <div className="message-input-area">
      <input
        type="text"
        placeholder="Message NOX..."
        value={message}
        disabled={isSending}
        onChange={(event) => onMessageChange(event.target.value)}
        onKeyDown={(event) => {
          if (event.key === 'Enter') {
            onSendMessage()
          }
        }}
      />

      <button
        className="send-button"
        disabled={isSending}
        onClick={onSendMessage}
      >
        {isSending ? '...' : '➤'}
      </button>
    </div>
  )
}

export default MessageInput