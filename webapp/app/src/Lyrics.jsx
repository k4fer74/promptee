import React, { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { render } from 'react-dom';
import './Lyrics.css';

const baseAPIPath = window.location.protocol + "//" + window.location.hostname + ":3160/api"

function Lyrics() {
    const [socketUrl, setSocketUrl] = useState('ws://'+ window.location.hostname + ':3160/api/prompter/broadcast/song_lyrics')
    const {sendMessage, lastMessage, readyState} = useWebSocket(socketUrl);
    const wsConnectionStatus = {
        [ReadyState.CONNECTING]: 'Connecting',
        [ReadyState.OPEN]: 'Open',
        [ReadyState.CLOSING]: 'Closing',
        [ReadyState.CLOSED]: 'Closed',
        [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
    }[readyState];

    const [songs, setSongs] = useState([]);

    const [songID, setSongID] = useState(0);
    const [song, setSong] = useState({
        lyrics: []
    });

    const [songLyrics, setSongLyrics] = useState({});

    useEffect(() => {
        let req = new XMLHttpRequest()
        req.open("GET", baseAPIPath + "/songs", false)
        req.send()
        setSongs(JSON.parse(req.responseText))
    }, []);

    useEffect(() => {
        if (songID !== 0) {
            let req = new XMLHttpRequest()
            req.open("GET", baseAPIPath + "/songs/" + songID, false)
            req.send()
            setSong(JSON.parse(req.responseText))
            setSongLyrics({
                i: 0,
                text: ""
            })
        }
    }, [songID])

    useEffect(() => {
        if (Object.keys(songLyrics).length !== 0) {
            const broadcastMessage = {
                author: song.author,
                song: song.name,
                text: songLyrics.text
            }
            sendMessage(JSON.stringify(broadcastMessage))
        }
    }, [songLyrics])

    function getSong(e) {
        setSongID(parseInt(e.target.value))
    }

    function stepLyrics(lyrics, i) {
        setSongLyrics({
           i: parseInt(i),
           text: lyrics,
        })
    }

    return (
        <div className="Lyrics">
            <header>
                <div>
                    <select onChange={getSong}>
                        <option>Song</option>
                        {songs.map(s => <option key={s.id} value={s.id}>{s.name} - {s.author}</option>)}
                    </select>
                </div>
            </header>
            <section>
                <ul>
                    {song.lyrics.map((l, i) =>
                        <li onClick={() => stepLyrics(l, i)} key={i} className={i === songLyrics.i ? "highlighted": ""}>
                            {l}
                        </li>
                    )}
                    <li>End of text</li>
                </ul>
            </section>
            <footer style={{display: songID === 0 ? "none": "block"}}>
                <ul>
                    <li onClick={() => stepLyrics(song.lyrics[songLyrics.i - 1], songLyrics.i - 1)} style={{display: songLyrics.i === 0 ? "none": "inline-block"}}>
                        Prev lyrics
                    </li>
                    <li onClick={() => stepLyrics(song.lyrics[songLyrics.i + 1], songLyrics.i + 1)} style={{display: songLyrics.i+1 !== song.lyrics.length ? "inline-block": "none"}}>
                        Next lyrics
                    </li>
                </ul>
            </footer>
        </div>
    )
}

render(<Lyrics />, document.getElementById('root'));