docker rm -f wiktionary-lemma;
docker rmi alaverydev/audio-language-wiktionary-lemma;
docker build -t alaverydev/audio-language-wiktionary-lemma .
docker push alaverydev/audio-language-wiktionary-lemma