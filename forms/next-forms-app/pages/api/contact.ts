// pages/api/contact.js
export default function handler(req, res) {
  if (req.method === 'POST') {
    const { email, subject, message } = req.body;

    // Do something with the data (e.g., save to a database, send an email, etc.)
    console.log({
      email,
      subject,
      message,
    });

    return res.status(200).json({ status: 'ok' });
  }

  return res.status(405).end(); // Method Not Allowed if not POST
}
