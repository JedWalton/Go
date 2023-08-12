import { useState, useRef, useEffect } from 'react';

const WebSocketPage = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [inputValue, setInputValue] = useState('');
  
  // Explicitly define the type for the ref as WebSocket | null
  const socketRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const socket = new WebSocket("ws://localhost:8080/echo");
    socketRef.current = socket;

    socket.onopen = () => {
      setMessages((prevMessages) => [...prevMessages, "Status: Connected"]);
    };

    socket.onmessage = (e: MessageEvent) => {
      setMessages((prevMessages) => [...prevMessages, "Server: " + e.data]);
    };

    return () => {
      socket.close();
    };
  }, []);

  const send = () => {
    if (socketRef.current) {
      socketRef.current.send(inputValue);
      setInputValue('');
    }
  };

  return (
    <div>
      <input 
        value={inputValue} 
        onChange={(e) => setInputValue(e.target.value)} 
        type="text" 
      />
      <button onClick={send}>Send</button>
      <pre>{messages.join('\n')}</pre>
    </div>
  );
}

export default WebSocketPage;

