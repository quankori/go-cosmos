FROM archwaynetwork/archwayd:constantine

COPY ./archwayd /.archway

RUN archwayd keys add cosmos-test --dry-run

CMD ["keys", "list"]