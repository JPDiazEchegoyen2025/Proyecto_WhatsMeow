// QRLogin.jsx
// Este componente consume el endpoint /login y muestra el QR
import React, { useEffect, useState } from 'react';
import QRCode from 'react-qr-code';

function QRLogin() {
  const [qr, setQr] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Hacemos la petición al backend
    fetch('http://localhost:8080/login')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Error al obtener el QR');
        }
        return response.json();
      })
      .then((data) => {
        // Suponemos que el backend responde con { qr: 'base64...' }
        setQr(data.qr);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Cargando QR...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!qr) return <p>No se recibió el QR.</p>;

  return (
    <div style={{ textAlign: 'center' }}>
      <h2>Escanea el QR para iniciar sesión en WhatsApp</h2>
      <QRCode value={qr} size={300} />
    </div>
  );
}

export default QRLogin;
