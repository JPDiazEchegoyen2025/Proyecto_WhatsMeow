import React, { useEffect, useState } from 'react';
import QRCode from 'react-qr-code';
// QRLogin.jsx
// Este componente consume el endpoint /login y muestra el QR

// Polling automático para detectar sesión activa tras escanear el QR
function QRLogin() {
  const [qr, setQr] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [logoutMsg, setLogoutMsg] = useState(null);
  const [showWelcome, setShowWelcome] = useState(true);
  const [qrTimer, setQrTimer] = useState(36);
  const [qrInterval, setQrInterval] = useState(null);
  const [attempts, setAttempts] = useState(0);
  const [sessionInfo, setSessionInfo] = useState(null); // Nueva variable para la sesión

  // Temporizador para el QR
  useEffect(() => {
    if (qr) {
      setQrTimer(36);
      if (qrInterval) clearInterval(qrInterval);
      const interval = setInterval(() => {
        setQrTimer((prev) => {
          if (prev <= 1) {
            clearInterval(interval);
            setQr(null);
            setError('El QR ha expirado.');
            return 0;
  useEffect(() => {
    let pollInterval = null;
    if (qr && !sessionInfo) {
      pollInterval = setInterval(() => {
        fetch('http://localhost:8080/session')
          .then((res) => res.json())
          .then((data) => {
            if (data && data.active) {
              setSessionInfo(data);
              setQr(null);
              setError(null);
              setLoading(false);
              clearInterval(pollInterval);
            }
          })
          .catch(() => {});
      }, 3000); // cada 3 segundos
    }
    return () => {
      if (pollInterval) clearInterval(pollInterval);
    };
  }, [qr, sessionInfo]);
          }
          return prev - 1;
        });
      }, 1000);
      setQrInterval(interval);
      return () => clearInterval(interval);
    }
    // Limpia el temporizador si no hay QR
    if (!qr && qrInterval) {
      clearInterval(qrInterval);
      setQrInterval(null);
    }
  }, [qr]);

  // Verificar sesión activa al cargar el componente
  useEffect(() => {
    fetch('http://localhost:8080/session')
      .then((res) => res.json())
      .then((data) => {
        if (data && data.active) {
          setSessionInfo(data);
          setShowWelcome(false);
          setLoading(false);
        } else {
          setSessionInfo(null);
          setShowWelcome(true);
          setLoading(false);
        }
      })
      .catch(() => {
        setSessionInfo(null);
        setShowWelcome(true);
        setLoading(false);
      });
  }, []);

  const handleLogout = async () => {
    setLogoutMsg(null);
    try {
      const res = await fetch('http://localhost:8080/logout', { method: 'POST' });
      const data = await res.json();
      if (res.ok) {
        setLogoutMsg(data.message);
        setQr(null);
        setError(null);
        setSessionInfo(null);
        setShowWelcome(true);
        setLoading(false);
      } else {
        setLogoutMsg(data.error || 'Error al cerrar sesión');
      }
    } catch (err) {
      setLogoutMsg('Error de red al cerrar sesión');
    }
  };

  // Función para volver a la pantalla de bienvenida y reiniciar estados
  const handleBack = () => {
    setShowWelcome(true);
    setQr(null);
    setError(null);
    setLoading(false);
    setLogoutMsg(null);
  };

  const handleStart = () => {
    setShowWelcome(false);
    setLoading(true);
    // Antes de iniciar el proceso de vinculación, volver a consultar la sesión
    fetch('http://localhost:8080/session')
      .then((res) => res.json())
      .then((data) => {
        if (data && data.active) {
          setSessionInfo(data);
          setShowWelcome(false);
          setLoading(false);
        } else {
          // Si no hay sesión, iniciar vinculación
          fetch('http://localhost:8080/login')
            .then((response) => {
              if (!response.ok) {
                setAttempts((prev) => prev + 1);
                throw new Error('Error al obtener el QR');
              }
              return response.json();
            })
            .then((data) => {
              setQr(data.qr);
              setLoading(false);
              setAttempts(0);
            })
            .catch((err) => {
              setError(err.message);
              setLoading(false);
            });
        }
      })
      .catch(() => {
        setError('Error al verificar la sesión');
        setLoading(false);
      });
  };

  // Mostrar información de la sesión si existe
  if (sessionInfo) return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <h2 style={{ color: '#202020', marginBottom: '16px' }}>¡Ya tienes una sesión activa!</h2>
        <p style={{ color: '#202020', fontSize: '1.1em', marginBottom: '16px' }}>
          Tu cuenta está vinculada y lista para usar.<br /><br />
          <b>Usuario:</b> {sessionInfo.user || 'Desconocido'}<br />
          <b>Teléfono:</b> {sessionInfo.phone || 'Desconocido'}<br />
          <b>Estado:</b> {sessionInfo.status || 'Activa'}
        </p>
        <p style={{ color: '#afafaf', fontSize: '1em', marginBottom: '16px' }}>
          Puedes entrar a la sesión activa para gestionar tus mensajes o cerrar la sesión si deseas vincular otra cuenta.
        </p>
        <button onClick={() => alert('Entrando a la sesión activa...')} style={{ marginTop: '10px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Entrar a la sesión</button>
        <button onClick={handleLogout} style={{ marginTop: '20px', background: '#d32f2f', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer', marginLeft: '10px' }}>Cerrar sesión</button>
      </div>
    </div>
  );
  // Si no hay sesión, mostrar bienvenida
  if (showWelcome) return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <h2 style={{ color: '#fff' }}>¡Bienvenido a WhatsApp Usando WhatsMeow!</h2>
      <p style={{ color: '#afafaf', fontSize: '1em', maxWidth: '500px' }}>
        WhatsMeow es una librería open source para interactuar con WhatsApp Web desde aplicaciones personalizadas, permitiendo automatizar, enviar y recibir mensajes, y gestionar sesiones de forma segura y flexible.
      </p>
      <p style={{ color: '#afafaf', fontSize: '1.1em' }}>Conecta tu cuenta de WhatsApp escaneando el código QR.</p>
      <button onClick={handleStart} className="start-btn">Comenzar</button>
    </div>
  );
  if (loading) return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <p style={{ color: '#202020', fontSize: '1.2em', marginBottom: '24px' }}><b>Cargando QR...</b></p>
        <p style={{ color: '#afafaf', fontSize: '1em', marginBottom: '16px' }}>Cuando aparezca el código QR, abre WhatsApp en tu teléfono, ve a <b>Menú &gt; Dispositivos vinculados</b> y escanea el código para conectar tu cuenta.</p>
        <button onClick={handleBack} style={{ marginTop: '20px' }}>Volver</button>
        {logoutMsg && <p>{logoutMsg}</p>}
      </div>
    </div>
  );
  if (error) return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <p style={{ color: '#202020', fontSize: '1.2em', marginBottom: '24px' }}><b>Error:</b> {error}</p>
        <button onClick={handleBack} style={{ marginTop: '20px', marginRight: '10px' }}>Volver</button>
        <button onClick={() => { setError(null); setLoading(true); setQr(null); handleStart(); }} style={{ marginTop: '20px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Actualizar QR</button>
        {attempts >= 5 && <p style={{ color: '#d32f2f', marginTop: '16px' }}>¿Problemas? Intenta más tarde o revisa tu conexión.</p>}
        {logoutMsg && <p>{logoutMsg}</p>}
      </div>
    </div>
  );
  if (!qr) return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <p style={{ color: '#202020', fontSize: '1.2em', marginBottom: '24px' }}><b>No se recibió el QR.</b></p>
        <button onClick={handleBack} style={{ marginTop: '20px' }}>Volver</button>
        {logoutMsg && <p>{logoutMsg}</p>}
      </div>
    </div>
  );

  return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto', display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
        <h2 style={{ color: '#202020', marginBottom: '16px' }}>Escanea el QR para iniciar sesión</h2>
        <div style={{ background: '#fff', padding: '24px', display: 'inline-block', borderRadius: '12px', boxShadow: '0 0 8px #eee', marginBottom: '16px' }}>
          <QRCode value={qr} size={300} />
        </div>
        <p style={{ color: '#afafaf', fontSize: '1em', marginBottom: '8px', textAlign: 'center' }}>
          Abre WhatsApp en tu teléfono, ve a <b>Menú &gt; Dispositivos vinculados</b> y escanea el código QR para conectar tu cuenta.
        </p>
        <p style={{ color: '#25d366', fontWeight: 'bold', marginBottom: '8px' }}>El QR expira en: {qrTimer}s</p>
        {qrTimer === 0 && (
          <button onClick={() => { setError(null); setLoading(true); setQr(null); handleStart(); }} style={{ marginTop: '10px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Actualizar QR</button>
        )}
        <button onClick={handleBack} style={{ marginTop: '20px' }}>Volver</button>
        {logoutMsg && <p>{logoutMsg}</p>}
      </div>
    </div>
  );
}

export default QRLogin;
