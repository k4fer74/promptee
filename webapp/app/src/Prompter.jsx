import React, { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { render } from 'react-dom';
import './Prompter.css';

function Prompter() {
    const [socketUrl, setSocketUrl] = useState('ws://' + window.location.hostname + ':3160/api/prompter/broadcast/prompt')
    const [messageHistory, setMessageHistory] = useState([]);
    const {sendMessage, lastJsonMessage, readyState} = useWebSocket(socketUrl, {
        onMessage: () => {
            console.log('received')
        }
    });
    useEffect(() => {
        if (lastJsonMessage !== null) {
            setMessageHistory((prev) => prev.concat(lastJsonMessage));
        }
    }, [lastJsonMessage, setMessageHistory]);
    const wsConnectionStatus = {
        [ReadyState.CONNECTING]: 'Connecting',
        [ReadyState.OPEN]: 'Open',
        [ReadyState.CLOSING]: 'Closing',
        [ReadyState.CLOSED]: 'Closed',
        [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
    }[readyState];

    return (
        <div className="Prompter">
            {lastJsonMessage &&
                <div className="BibleText">
                    <h1>{lastJsonMessage.book_name} {lastJsonMessage.chapter_number}:{lastJsonMessage.verse_number}</h1>
                    <h2>{lastJsonMessage.verse_text}</h2>
                </div>
            }
        </div>
    )
}

render(<Prompter />, document.getElementById('root'));