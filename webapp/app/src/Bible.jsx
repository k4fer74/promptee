import React, { useState, useEffect } from 'react';
import useWebSocket, { ReadyState } from 'react-use-websocket';
import { render } from 'react-dom';
import './Bible.css';

const baseAPIPath = window.location.protocol + "//" + window.location.hostname + ":3160/api"

function Bible() {
    const [socketUrl, setSocketUrl] = useState('ws://'+ window.location.hostname + ':3160/api/prompter/broadcast/bible_text')
    const {sendMessage, lastMessage, readyState} = useWebSocket(socketUrl);
    const wsConnectionStatus = {
        [ReadyState.CONNECTING]: 'Connecting',
        [ReadyState.OPEN]: 'Open',
        [ReadyState.CLOSING]: 'Closing',
        [ReadyState.CLOSED]: 'Closed',
        [ReadyState.UNINSTANTIATED]: 'Uninstantiated',
    }[readyState];

    const [books, setBooks] = useState([]);
    const [bookNumber, setBookNumber] = useState(0);

    const [chapters, setChapters] = useState([]);
    const [chapterNumber, setChapterNumber] = useState(0);

    const [verses, setVerses] = useState([]);
    const [verseNumber, setVerseNumber] = useState(0);

    useEffect(() => {
        let req = new XMLHttpRequest()
        req.open("GET", baseAPIPath + "/books", false)
        req.send()
        setBooks(JSON.parse(req.responseText))
    }, []);

    useEffect(() => {
        if (bookNumber !== 0) {
            let req = new XMLHttpRequest()
            req.open("GET", baseAPIPath + "/books/" + bookNumber + "/chapters", false)
            req.send()
            setChapters(JSON.parse(req.responseText))
            setChapterNumber(0)
            setVerseNumber(0)
            setVerses([])
        }
    }, [bookNumber])

    useEffect(() => {
        if (chapterNumber !== 0) {
            let req = new XMLHttpRequest()
            req.open("GET", baseAPIPath + "/books/" + bookNumber + "/chapters/" + chapterNumber + "/verses", false)
            req.send()
            setVerseNumber(1)
            setVerses(JSON.parse(req.responseText))
        }
    }, [chapterNumber])

    useEffect(() => {
        if (verseNumber !== 0) {
            const broadcastMessage = {
                book_number: bookNumber,
                book_name: books[bookNumber-1].name,
                chapter_number: chapterNumber,
                verse_number: verseNumber,
                verse_text: verses[verseNumber-1].text
            }
            sendMessage(JSON.stringify(broadcastMessage))
        }
    }, [verseNumber])

    function getChapters(e) {
        setBookNumber(parseInt(e.target.value))
    }

    function getVerses(e) {
        setChapterNumber(parseInt(e.target.value))
    }

    function jumpToSingleVerse(e) {
        setVerseNumber(parseInt(e.target.value))
    }

    return (
        <div className="Bible">
            <header>
                <div>
                    <select onChange={getChapters}>
                        <option>Book</option>
                        {books.map(b => <option key={b.number} value={b.number}>{b.name}</option>)}
                    </select>
                </div>
                <div>
                    <select value={chapterNumber} onChange={getVerses}>
                        <option>Chapter</option>
                        {chapters.map(c => <option key={c.number} value={c.number}>{c.number}</option>)}
                    </select>
                </div>
                <div>
                    <select value={verseNumber} onChange={jumpToSingleVerse}>
                        <option>Jump to</option>
                        {verses.map(v => <option key={v.number} value={v.number}>{v.number}</option>)}
                    </select>
                </div>
            </header>
            <section>
                <ul>
                    {verses.map(v =>
                        <li onClick={() => setVerseNumber(v.number)} key={v.number} className={v.number === verseNumber ? "highlighted": ""}>
                            <span>{v.number}</span> {v.text}
                        </li>
                    )}
                    <li>End of text</li>
                </ul>
            </section>
            <footer style={{display: chapterNumber === 0 ? "none": "block"}}>
                <ul>
                    <li onClick={() => setVerseNumber(verseNumber-1)} style={{display: verseNumber === 1 ? "none": "inline-block"}}>
                        Prev verse ({verseNumber-1})
                    </li>
                    <li onClick={() => setVerseNumber(verseNumber+1)} style={{display: verseNumber !== verses.length ? "inline-block": "none"}}>
                        Next verse ({verseNumber+1})
                    </li>
                </ul>
            </footer>
        </div>
    )
}

render(<Bible />, document.getElementById('root'));