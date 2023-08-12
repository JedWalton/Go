// pages/index.js
import { useState } from 'react';

export default function Home() {
  const [success, setSuccess] = useState(false);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const formData = {
      email: e.currentTarget.email.value,
      subject: e.currentTarget.subject.value,
      message: e.currentTarget.message.value,
    };

    const response = await fetch('http://localhost:8080/', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(formData),
    });

    if (response.status === 200) {
      setSuccess(true);
    }
  };

  return (
    <div>
      {success ? (
        <h1>Thanks for your message!</h1>
      ) : (
        <>
          <h1>Contact</h1>
          <form onSubmit={handleSubmit}>
            <label>Email:</label>
            <br />
            <input type="text" name="email" />
            <br />
            <label>Subject:</label>
            <br />
            <input type="text" name="subject" />
            <br />
            <label>Message:</label>
            <br />
            <textarea name="message"></textarea>
            <br />
            <input type="submit" value="Submit" />
          </form>
        </>
      )}
    </div>
  );
}

