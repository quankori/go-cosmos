FROM archwaynetwork/archwayd:constantine

COPY ./archwayd /root/.archway/config
# RUN archwayd keys add cosmos-test --dry-run
RUN ls -la /root/.archway/config
CMD ["start", "--x-crisis-skip-assert-invariants"]