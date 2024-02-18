INSERT INTO
    books (title, author_id)
VALUES
    (
        'Hamlet',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'Macbeth',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'Othello',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'Romeo and Juliet',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'King Lear',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'Julius Caesar',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'The Tempest',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Shakespeare'
        )
    ),
    (
        'War and Peace',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Tolstoy'
        )
    ),
    (
        'Crime and Punishment',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Dostoevsky'
        )
    ),
    (
        'Anna Karenina',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Tolstoy'
        )
    ),
    (
        'The Brothers Karamazov',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Dostoevsky'
        )
    ),
    (
        'The Idiot',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Dostoevsky'
        )
    ),
    (
        'A Farewell to Arms',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Hemingway'
        )
    ),
    (
        'The Great Gatsby',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Fitzgerald'
        )
    ),
    (
        'Tender Is the Night',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Fitzgerald'
        )
    ),
    (
        'This Side of Paradise',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Fitzgerald'
        )
    ),
    (
        '1984',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Orwell'
        )
    ),
    (
        'Animal Farm',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Orwell'
        )
    ),
    (
        'Brave New World',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Huxley'
        )
    ),
    (
        'The Time Machine',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Wells'
        )
    ),
    (
        'The War of the Worlds',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Wells'
        )
    ),
    (
        'Twenty Thousand Leagues Under the Sea',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Verne'
        )
    ),
    (
        'Les Misérables',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Hugo'
        )
    ),
    (
        'The Count of Monte Cristo',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Dumas'
        )
    ),
    (
        'The Three Musketeers',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Dumas'
        )
    ),
    (
        'The Adventures of Tom Sawyer',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Twain'
        )
    ),
    (
        'The Adventures of Huckleberry Finn',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Twain'
        )
    ),
    (
        'The Adventures of Sherlock Holmes',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Doyle'
        )
    ),
    (
        'The Moonstone',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Collins'
        )
    ),
    (
        'Lost city',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Call of Cthulhu',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Shadow over Innsmouth',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Dunwich Horror',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Colour Out of Space',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Haunter of the Dark',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Outsider',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'At the Mountains of Madness',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Whisperer in Darkness',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Case of Charles Dexter Ward',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Dreams in the Witch House',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Thing on the Doorstep',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'The Shadow out of Time',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Lovecraft'
        )
    ),
    (
        'Pride and Prejudice',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Sense and Sensibility',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Emma',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Mansfield Park',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Northanger Abbey',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Persuasion',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Austen'
        )
    ),
    (
        'Jane Eyre',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Bronte'
        )
    ),
    (
        'Wuthering Heights',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Bronte'
        )
    ),
    (
        'Agnes Grey',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Bronte'
        )
    ),
    (
        'The Tenant of Wildfell Hall',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Bronte'
        )
    ),
    (
        'Mrs. Dalloway',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Woolf'
        )
    ),
    (
        'To the Lighthouse',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Woolf'
        )
    ),
    (
        'Orlando',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Woolf'
        )
    ),
    (
        'The Waves',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Woolf'
        )
    ),
    (
        'Middlemarch',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Mill on the Floss',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Silas Marner',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Daniel Deronda',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Adam Bede',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Felix Holt, the Radical',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Romola',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Spanish Gypsy',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Legend of Jubal',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Brother Jacob',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Lifted Veil',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Mr. Gilfil''s Love Story',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'Janet''s Repentance',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Black Monk',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    ),
    (
        'The Sad Fortunes of the Reverend Amos Barton',
        (
            SELECT
                id
            FROM
                authors
            WHERE
                name = 'Eliot'
        )
    );
