import React, { useEffect, useState, useRef } from 'react';
import QRCode from 'react-qr-code';
import ChatPanel from './pages/ChatPanel';
// QRLogin.jsx

function QRLogin() {
  // Estados principales de la vista
  const [view, setView] = useState('loading'); // 'loading' | 'welcome' | 'qr' | 'session' | 'error'
  const [qr, setQr] = useState(null);
  const [qrTimer, setQrTimer] = useState(36);
  const [sessionInfo, setSessionInfo] = useState(null);
  const [error, setError] = useState(null);
  const [logoutMsg, setLogoutMsg] = useState(null);
  const [attempts, setAttempts] = useState(0);
  const qrIntervalRef = useRef(null);
  const pollIntervalRef = useRef(null);

  // Verificar sesión activa al cargar
  useEffect(() => {
    fetch('http://localhost:8080/session')
      .then((res) => res.json())
      .then((data) => {
        if (data && data.active) {
          setSessionInfo(data);
          setView('session');
        } else {
          setSessionInfo(null);
          setView('welcome');
        }
      })
      .catch(() => {
        setSessionInfo(null);
        setView('welcome');
      });
  }, []);

  // Temporizador para el QR
  useEffect(() => {
    if (view === 'qr' && qr) {
      setQrTimer(36);
      if (qrIntervalRef.current) clearInterval(qrIntervalRef.current);
      const interval = setInterval(() => {
        setQrTimer((prev) => {
          if (prev <= 1) {
            clearInterval(interval);
            setQr(null);
            setError('El QR ha expirado.');
            setView('error');
            return 0;
          }
          return prev - 1;
        });
      }, 1000);
      qrIntervalRef.current = interval;
      return () => clearInterval(interval);
    }
    // Limpia el temporizador si no hay QR
    if (view !== 'qr' && qrIntervalRef.current) {
      clearInterval(qrIntervalRef.current);
      qrIntervalRef.current = null;
    }
  }, [view, qr]);

  // Polling para detectar sesión activa tras escanear el QR
  useEffect(() => {
    if (view === 'qr' && qr && !sessionInfo) {
      if (pollIntervalRef.current) clearInterval(pollIntervalRef.current);
      const poll = setInterval(() => {
        fetch('http://localhost:8080/session')
          .then((res) => res.json())
          .then((data) => {
            if (data && data.active) {
              setSessionInfo(data);
              setQr(null);
              setError(null);
              setView('session');
              clearInterval(poll);
            }
          })
          .catch(() => {});
      }, 3000);
      pollIntervalRef.current = poll;
      return () => clearInterval(poll);
    }
    if (view !== 'qr' && pollIntervalRef.current) {
      clearInterval(pollIntervalRef.current);
      pollIntervalRef.current = null;
    }
  }, [view, qr, sessionInfo]);

  // Iniciar proceso de vinculación
  const handleStart = () => {
    setView('loading');
    fetch('http://localhost:8080/session')
      .then((res) => res.json())
      .then((data) => {
        if (data && data.active) {
          setSessionInfo(data);
          setView('session');
        } else {
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
              setAttempts(0);
              setView('qr');
            })
            .catch((err) => {
              setError(err.message);
              setView('error');
            });
        }
      })
      .catch(() => {
        setError('Error al verificar la sesión');
        setView('error');
      });
  };

  // Volver a bienvenida y limpiar estados
  const handleBack = () => {
    setQr(null);
    setError(null);
    setLogoutMsg(null);
    setAttempts(0);
    setView('welcome');
  };

  // Cerrar sesión
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
        setView('welcome');
      } else {
        setLogoutMsg(data.error || 'Error al cerrar sesión');
      }
    } catch (err) {
      setLogoutMsg('Error de red al cerrar sesión');
    }
  };

  // Renderizado por estado
  if (view === 'loading') return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <p style={{ color: '#202020', fontSize: '1.2em', marginBottom: '24px' }}><b>Cargando...</b></p>
      </div>
    </div>
  );

  if (view === 'session' && sessionInfo) return (
    <div style={{ width: '100%' }}>
      <ChatPanel user={sessionInfo.user} onLogout={handleLogout} />
      {/* Aquí puedes agregar el panel de chat o más contenido */}
      {console.log('sessionInfo:', sessionInfo)}
      {(logoutMsg || sessionInfo) && (
        <div style={{ position: 'fixed', left: 0, bottom: 0, width: '100vw', background: '#202020', color: '#d32f2f', display: 'flex', alignItems: 'center', justifyContent: 'space-between', padding: '2px 0 3px 0', fontWeight: 'normal', fontSize: '0.95em', zIndex: 200, borderTop: '1px solid #888' }}>
          <span style={{ marginLeft: 12, fontWeight: 'bold', letterSpacing: 1 }}>&gt; {logoutMsg || 'Sesión activa'}</span>
          <span style={{ display: 'flex', alignItems: 'center', marginRight: 12 }}>
            <span style={{ color: '#888', margin: '0 12px' }}>|</span>
            <span style={{ marginRight: 18 }}>Móvil: <span style={{ color: '#fff' }}>{sessionInfo?.user ? `+${sessionInfo.user.replace(/:.*$/, '')}` : 'N/A'}</span></span>
            <span style={{ color: '#888', margin: '0 12px' }}>|</span>
            <span>Estado: <span style={{ color: '#25d366' }}>{sessionInfo?.active ? 'Activa' : 'Inactiva'}</span></span>
          </span>
        </div>
      )}
    </div>
  );

  if (view === 'welcome') return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <h2 style={{ color: '#fff' }}>¡Bienvenido a WhatsApp Usando WhatsMeow!</h2>
      <p style={{ color: '#afafaf', fontSize: '1em', maxWidth: '500px' }}>
        WhatsMeow es una librería open source para interactuar con WhatsApp Web desde aplicaciones personalizadas, permitiendo automatizar, enviar y recibir mensajes, y gestionar sesiones de forma segura y flexible.
      </p>
      <p style={{ color: '#afafaf', fontSize: '1.1em' }}>Conecta tu cuenta de WhatsApp escaneando el código QR.</p>
      <button onClick={handleStart} className="start-btn">Comenzar</button>
    </div>
  );

  if (view === 'qr' && qr) return (
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
        <button onClick={handleBack} style={{ marginTop: '20px' }}>Volver</button>
      </div>
    </div>
  );

  if (view === 'error') return (
    <div style={{ textAlign: 'center', minHeight: '60vh', display: 'flex', flexDirection: 'column', alignItems: 'center', justifyContent: 'center' }}>
      <div style={{ background: '#fff', padding: '40px 32px', borderRadius: '16px', boxShadow: '0 0 16px #222', maxWidth: '400px', margin: '0 auto' }}>
        <p style={{ color: '#202020', fontSize: '1.2em', marginBottom: '24px' }}><b>Error:</b> {error}</p>
        <button onClick={handleBack} style={{ marginTop: '20px', marginRight: '10px' }}>Volver</button>
        <button onClick={() => { setError(null); setView('loading'); handleStart(); }} style={{ marginTop: '20px', background: '#25d366', color: '#fff', border: 'none', borderRadius: '8px', padding: '10px 24px', fontWeight: 'bold', cursor: 'pointer' }}>Actualizar QR</button>
        {attempts >= 5 && <p style={{ color: '#d32f2f', marginTop: '16px' }}>¿Problemas? Intenta más tarde o revisa tu conexión.</p>}
      </div>
      {logoutMsg && (
        <div style={{ position: 'fixed', left: 0, bottom: 0, width: '100vw', background: '#25d366', color: '#fff', textAlign: 'center', padding: '12px 0', fontWeight: 'bold', fontSize: '1.1em', zIndex: 200 }}>
          {logoutMsg}
        </div>
      )}
    </div>
  );

  return null;
}

export default QRLogin;

