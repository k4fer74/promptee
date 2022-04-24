import React, { useState, useEffect } from 'react';
import { render } from 'react-dom';
import './Bible.css';

function Bible() {
    const [books, setBooks] = useState([]);
    const [bookNumber, setBookNumber] = useState(0);

    const [chapters, setChapters] = useState([]);
    const [chapterNumber, setChapterNumber] = useState(0);

    const [verses, setVerses] = useState([]);
    const [verseNumber, setVerseNumber] = useState(0);

    useEffect(() => {
        let req = new XMLHttpRequest()
        req.open("GET", "http://127.0.0.1:3160/api/books", false)
        req.send()
        setBooks(JSON.parse(req.responseText))
    }, []);

    useEffect(() => {
        if (bookNumber !== 0) {
            let req = new XMLHttpRequest()
            req.open("GET", "http://127.0.0.1:3160/api/books/" + bookNumber + "/chapters", false)
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
            req.open("GET", "http://127.0.0.1:3160/api/books/" + bookNumber + "/chapters/" + chapterNumber + "/verses", false)
            req.send()
            setVerseNumber(1)
            setVerses(JSON.parse(req.responseText))
        }
    }, [chapterNumber])

    function getChapters(e) {
        setBookNumber(parseInt(e.target.value))
    }

    function getVerses(e) {
        setChapterNumber(parseInt(e.target.value))
    }

    function getSingleVerse(e) {
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
                    <select value={verseNumber} onChange={getSingleVerse}>
                        <option>Jump to</option>
                        {verses.map(v => <option key={v.number} value={v.number}>{v.number}</option>)}
                    </select>
                </div>
            </header>
            <section>
                <ul>
                    {verses.map(v =>
                        <li key={v.number} className={v.number === verseNumber ? "highlighted": ""}>
                            <span>{v.number}</span> {v.text}
                        </li>
                    )}
                </ul>
            </section>
        </div>
    )
}


render(<Bible />, document.getElementById('root'));