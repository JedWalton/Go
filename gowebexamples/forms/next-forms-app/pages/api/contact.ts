// pages/api/contact.js
export default function handler(
  req: {
    method: string; body: { 
      email: any;
      subject: any;
      message: any; };
  }, res: {
    status: (arg0: number) => { (): any; new(): any;
      json: { (arg0: { status: string; }): any; new(): any; }; end:
      { (): any; new(): any; }; }; }): any {
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
