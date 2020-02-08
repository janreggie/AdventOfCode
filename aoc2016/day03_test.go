package aoc2016

import (
	"testing"

	"github.com/janreggie/AdventOfCode/tools"
	"github.com/stretchr/testify/assert"
)

const day03myInput = `  810  679   10
783  255  616
545  626  626
 84  910  149
607  425  901
556  616  883
938  900  621
638  749  188
981  415  634
680  557  571
523  604  270
910  954  484
464  392  514
458   52  687
696  438  832
213  583  966
572  571  922
451   42  686
177  390  688
151  136  705
 92  413  191
789  676  377
486  262  600
450  708  472
556    9  481
157   85   94
574   93  549
539  165  487
815  742   73
353  773  428
526  152  680
433  711  557
168  632  306
848  992  757
885  786  890
469  475  146
899  833  137
864  202  688
101  902  620
529  937  826
 41  381  521
562  883  804
468  197  272
451    8  420
561  193  630
597  951  383
171  845  251
541  810  157
268   46  712
332    2  397
100   47  436
194  665  205
325  277   21
170  652  205
765  165  506
 15  257  144
762  124  401
662  543  531
 29  425  308
667  785  299
935  758  405
504  998  367
771  947  630
490  933  978
441  498  896
862  896  607
655  935  194
286  240  324
368  723  311
419  762  600
316  903  529
197  215  215
551  461   77
855  318    7
894  690   86
451  648  416
608  132  385
420  761  112
560  711  195
371  750  506
188  307  584
 26  377  622
304  701  292
286  630  642
883  880  379
774  564  597
300  692  701
529  595   27
740   76  445
567  648  422
340  163  901
374  775  902
308  827  882
529  371  374
996  587  162
534  360  516
924  160  276
724  896  687
929  971  578
798  252  761
512  991  812
465  758   49
724  446  571
482  196  544
553  247   86
624  552  778
 73  143  127
556  471  749
224  927  383
133  636  847
174  985  569
572  819  881
282  818  383
535  429  780
953  540  815
577  302  494
530  654  370
670  739  168
700  695  806
196   48  928
255  805  749
 65   96  969
292  860  929
556  269  297
 43  832  407
542  723  438
919  139  407
709  194  955
847  237  933
321   41  216
778  749  374
782  745  529
716  572  251
 90   49  976
639  557  740
148  125  784
143  819  382
 71  729  563
309  500  806
 25  412  594
296  600  237
681  187  142
758  913  288
163  972  266
197  352  190
383  190  562
206  214  393
566  307  294
  2  284  335
564  472  394
635  928  589
169  744  574
710  386  589
970  386  827
943  424  134
846  269  712
266  765  615
344  824  685
250  222  554
377  586  859
398  526  275
317  996  937
503  364  389
212  782  533
584  539  589
731  200  584
773  389  578
 43  482  104
432  140  339
193  758  673
612  882  582
314  920  130
522   40   26
695  939  149
955  121  552
728  850  661
524  766  433
817  221  992
753  580  543
 72  392  873
445  897    3
144  508  567
354  990  566
477  392  687
602  846  520
321  577  677
716  518   55
367   77  545
361  473  504
 98  893  887
854  920  887
860  174   30
389  857  797
686  968  907
613  275  595
855  440  906
749  494  735
527  895  550
767  971  488
118  814  148
854  193  480
847  425  378
697  159  357
282  476   48
 96  314  176
949  597  903
956  478  885
714  754  278
757  547  210
 53  223  170
355  725  928
930  780  762
924  581  266
570  132  283
625  674  529
159  719  325
316  670  929
 55  655  542
344   19  791
437  805  312
327  867  647
521  405  496
383   58  117
638   36  175
924   59  112
401   66  353
740  785  823
713  725  622
821  702  246
378   24  958
690  718  924
486  788  537
377  214  670
514  720  427
451  927  877
808  868  872
554   94    2
534  516  715
735  318  125
880  496  755
724  115  567
 23  105   89
725   55  561
599   44  581
378  661  173
628  640  632
747  817  448
557  248  338
743  833  776
309  895  759
 18  696  851
328  775  356
220   37  499
865  390  651
736  397  205
645  949  170
638  860  143
 23  262   98
822   46  842
663  687  860
941  700  745
762  304  509
154  275  369
728  155  324
 99  113  485
245   82   62
294   76  484
215  664  398
146  336  461
102  591  503
535  814  749
250  410  892
672  467  212
304  108  285
300  246   11
  4  304  284
115  132  112
460  334  739
453  281  792
505  591    6
482  413  975
 26  763  980
226  377  727
406   59   39
570  325  691
333  438  966
267  792  229
130  384  854
375  165  187
 37  498  403
357  509  242
710  796  296
708  187  265
 46  762  279
 84  589  760
578   38  226
624  558  570
338  517  276
547  498  648
626  265  677
144  662  193
581  820  407
477  567  232
582  890  926
167  458  502
635  841  607
505  346  239
522  970  506
608  830  686
100   89  353
 95  159  652
 24  163  786
328  313  534
793   52  249
750  274  683
885  463  247
534  326  391
938  726  199
893  620  120
899  410  508
226  896  459
677  694  780
880   15  831
909  683  903
 55    7  541
294  221  109
286  216  507
239  652  380
948  760  431
772  258  275
562  226  631
503  264  765
690   42  369
761  541  373
232  596   75
925   60  402
550  181   16
600  579  701
 92  419  696
 26  117  290
  4  487  157
 21  474  308
 99  827  835
279  216  451
267  739  749
309  456  262
320   91  282
 52  431  304
773  784  932
474  483  932
703  975  257
851  227  584
 17  224  365
845   96  536
258  150  905
797  119  876
862  196  220
954  964  355
534  979  302
905  509  628
153  185  273
169  538  509
 43  477  356
702  357  940
340  403  284
638   86  744
329  426  903
222  720  682
127  624  253
 28  849  485
555  158  599
553  690  443
598  926  185
611  934  868
986    8  983
166  396  946
500  822  662
507  715  828
294  790  587
661  779  235
549  594  657
771  918  800
923  896  983
866  203  437
723  465  852
589  717  731
332  331  710
984  484  794
750  479  886
857    5  286
400  841   63
665  513  508
841  739  513
331  586  669
420  561  690
346  104   22
847  758  149
570  211  816
524  868  962
483  229  317
408  555  325
682  650  285
646  987  974
467  368  779
442  640  968
644  131  184
903  916  162
565  890   91
474  763  351
569  178  709
520  618  666
437   75  213
509  471  758
298  486  904
364  416  429
513  971  271
169  863  202
 15  206  565
163   69  713
167  186  542
908  550   89
936  764  451
118  467  464
 89  385  375
179  165  545
143  514  187
313   47  636
477  830  550
769  808  577
 74  756  630
698  799  654
721  387   36
993  763  945
707  746    7
955  113  948
723  532  526
174  795  204
671  968  575
523  256  109
570  186  296
350  351  215
141  251   22
532  217  695
460   37  719
695   69  516
 36  597  350
670  552  556
287  143   35
400  801   45
133  921   71
637  169  646
108  721  890
655  681  311
885  393  603
375  388  113
976  522  534
 15  516  627
685  602  535
669  390  781
845  950  348
388   30  379
825  955   46
360  579  898
363  573  660
 33   30  864
905  723  916
968  648  655
178  181  363
754  262  268
883  837   45
216  687  222
520  973  909
808  968  943
335    3  202
211  605  517
 32  298  358
184  488  173
741   23  328
400  482  144
626  491  451
920  546  219
363  734  861
739  417  685
954  470  541
598  679  950
550  372  450
980  459  213
353  374  293
720  220  256
173   29  571
289  769  833
372  793  345
578  298  332
763  225  167
258  519  307
504    7  649
186  319  883
358  322  918
293   60  330
373  562  550
310  532  573
741  129  533
701  614  869
 54  736  587
451  131  817
499  784  651
931  681  193
674  311  500
900  312  197
553   94  331
  9  715  572
590   97  275
579  713  299
 20  345  741
817  738  534
819  963  497
168  303  997
462  599  698
400  772  485
755  922  928
591  847  180
500  135  977
946  940  751
658  368  790
720  714  141
850  261  594
615  116  476
660  156  488
485  895  378
797  992  614
847  652  838
842  516  364
745  444  329
175  362   84
684  223  578
 43  291  394
702  222  862
208  247  494
601  236  234
780   53  675
754  135  126
 26  776   52
735  716  136
591  829  171
606  373  824
 51  926  766
273  161  558
215  557  149
393  703  653
318  208  207
891   54  570
790  153  689
521  693  423
559  986  542
 58  611  404
178  509  602
684  120  975
791  407  811
 94  321   66
 14  317  266
108   14  271
580  454  391
781   82  849
419  406  775
396  298  237
448  375  330
747  301  322
103  835  120
138  897  630
127  102  546
518  552  412
398  442   43
586  972  380
 30  535   91
 42  384  962
 61  414  942
610  147   65
945  155  418
667   54  375
473  251  187
440  222  124
886  158  163
862  493  149
805  451  536
 59  108  458
663  613  719
264  525  574
755  176  168
390    6  783
 50  561  233
401  568  582
121  979  769
 94   77  830
195  938  201
124  626  161
668  633   35
662   29  164
394  658  768
203  918  850
466  425  399
353  804  714
323  851  640
152  939  642
 29  309  484
579  529  822
608  262  731
 38  756  450
433  828  740
431  895  693
392  477  399
 25  925  513
368  969  491
671  736  911
307  198  660
662  859  311
853  596  526
917   24  461
677  574  960
697  220   90
203  458  102
499  284   29
400   79  582
484  195  597
575  276  912
493  269  347
 23  593  223
476  802  358
 33  944  255
715  117  460
739  885  586
748  954  527
734  773  643
542  202  117
 15  976  460
309  830  331
319  208  557
458  822  461
545  784  690
878  372  858
 57  295  470
268  537  822
271  301  699
806  909  878
744  182  571
106  895  468
121  778   28
641  202  593
710  724  592
125  784  603
654  771   83
721   87  543
585  724   89
381  739  524
623   28  494
869  729  292
228  736  298
803   10   95
700  224  786
738  512    9
708  407  775
558  645  863
 45  209  466
540  809  587
372  512  717
416  203  974
272  496  928
816  141  903
675  894   84
567  900  957
827  122  189
882  860   56
 98  792  196
861  461  209
685  339   87
585  464  235
640  156  703
817  596  321
893  462  996
679  536  208
199  455  365
873  260  492
528  179  563
689  563  849
887  417  507
 64  270  198
595  214  166
566  232  242
921  102  212
187  202  335
992  169  475
736  754  200
655  374  127
 84  492  193
 21  709  972
199  208  236
216  683  926
479  669  604
437  872  293
789  256  515
341  948  637
142  933  536
207   82  218
702  249  779
253  369  874
508  255  254
 91  536  541
212  813   28
144  406  563
180  513  277
421  842  639
570  520  522
224  830  592
153  582  606
 81  415  239
160  553  735
525  348  778
454  352  626
609  460  169
559   57  334
784  428  242
706  867  289
637  914  281
620  407   83
152  446   90
260  331  799
301  677  725
708  254  328
418  147  798
732  344  963
627  626  302
670  241   76
220  383  376
733  124   50
795  673  466
136  637  423
823  258  700
204  936  878
730  976  981
272  310  894
333  201  863
 90  122  621
 90  811  209
275  904  283
193  125  189
127  961  283
347  529  829
352  738  734
878  726  411
942   54   34
429  750  426
367  938  424
501  447  757
566  773  648
382  140  899
462  353   90
230  493  945
425  290  415
894  360   21
897  529  431
914  124  338
 78  766  876
858  664  764
598  664  317
630  548  772
 30  483  604
642  331  545
518  702  474
546  750  887
252  663  547
813  917  671
852  367  894
 97  192  265
661  587  858
726  674  748
578  178  878
327  535  608
426  419  871
559  837  229
851  721  708
860  978  770
308  604  626
198  168  408
138  628  799
669  525  918
804  762  652
389  429  554
618  566  360
814  648  887
677  697  659
600  660  162
256  749  195
840  734  216
445  192  960
341  226  975
699  140  114
763  833  533
234  835   38
798   10  569
190  745  418
183  563  486
295  224  197
437  724  885
197  706  328
268  709  702
351  679  694
642  555  769
333  521  883
182  532  772
517  543  711
657  154  169
134  888  300
217  121  209
346  796  100
755  681  817
277  733  980
677  162  481
527  191  433
293  999  653
429  850  503
562  205  402
217  323  414
565  402   43
730  223  537
  4  701  567
737  570  523
644  510  459
390  252  367
344  715  179
 62  236  586
527  310  137
526   96  548
585  357  407
768  532  384
591  421   43
928  129  533
228  469  848
886  349  596
392  231  867
507  664  870
546  881  121
 28  306  275
688  284  261
683  495   31
733  191  899
 83  785  730
738  668  220
795   69  237
148  175  238
872  139  100
673  671  744
222  421  346
824  971  589
283  135  474
626   48  487
426  172  548
796  463  616
547  349  568
717  798  428
248  977  192
337  683  128
480  487  231
817  559  882
413  935  879
694  724  447
221  458  449
649  523  725
689  131  311
726  707  273
712  689  127
 65  338  183
612  523  679
631  834  297
701  320  433
265  518  602
691  519  160
463    4  575
777  590  394
790  975  201
 22  449  242
578  308  911
371  157  191
489  263  789
962  696  390
494  760  494
760  656  350
 57  322  551
639  105  616
676  402  236
269  464  893
265  573  312
472  822  682
410  385  584
882   56  493
596  330  827
184  494  873
 61  580  793
157  260  128
440  239  390
701  174  230
946  357  394
273  423  258
529  438  733
552   75  892
946  755  996
 64  836  112
971  192  928
188  378  692
179  299  676
 91  177  202
748  644  634
551  355  345
265  504  410
644   58  450
103  716  556
691  679  128
166  255  174
415  682  368
474  862  434
348  462  133
704  626  374
979  835  426
239  897  288
381  953  234
181   65  504
 61  803  297
761   22  946
771  822  908
900  914  563
656  948  114
349  202  594
322  294  811
535  484  837
532  438  869
700   94  814
691  557  159
201  512  738
598  652  742
269  642  772
698   23   49
376  375  689
375  476  819
426  421  559
683  775  420
876  374  995
281  556  587
990  137  273
782  928  299
895  829   65
228  687  764
 62  496  905
210  277  352
732  461  535
418  364  561
958  373  189
640  617   27
185  680  698
697  507  688
324  836  143
434  868  658
342  516  628
351  760  280
796  663  876
977  133  813
169  326  101
139  575  796
236  597  851
191  704  375
568  733  436
615   68  728
478  768  617
531  594  596
898  898   64
596  181  707
371  381  259
609  406  528
810  271  308
211  975  596
963  896  551
 94  362  418
812  351  848
732  495  708
866  246  209
973  682  792
898  535  672
667  237  783
325  642  229
419  654  754
328  374    7
359  468   93
 91  453   93
923  741   53
721  938  589
235  716  605
466  387  199
554  430  681
166  181  864
699  998  953
999  962  718
330  124  822
443  536  930
293  631  674
197  574  315
407  183  293
432  417  537
 31  571  657
901  555  463
686  456  465
217  259    3
742  535  427
881  347  555
769  659  299
134  577   20
252  566  877
181   10  885
191  829  994
744  649  867
910  354  781
 68  767  930
 88  716  850
 22  290  121
226  212  666
266  327  812
356  112  148
252  397  741
325  674  834
389  442  946
898   83  618
 51  807  862
844  772  461
831  546  467
644  476  539
758  758  722
346  512  463
157  427  697
439  672  243
192  869  150
890  977  753
962  767  607
818  926  500
960  927  219
377    9  389
661  191  869
695  149  368
358  342  778
474  396  202
546  585  853
 74  281  734
830  295  611
 19  813  388
847  963  378
 78  140  278
531  580  246
550  546  415
739  419  197
803  266  247
285  672  123
669   51  665
525  662    5
998  619  667
737  368  910
533  550  245
899  667  932
 80  302  566
508    1  576
454  303   15
752  463  159
119  380  906
702  279  942
234  198  326
262  207  305
214  388   64
975  779  523
975  243  519
694  895   79
750  477  112
746  470  108
201  299  119
748  890  652
808  897  387
908  617  466
739  750  302
887  765  558
464   97  662
 11  745  109
454  537   27
446  363  118
265   33  670
862  497  147
681  488  582
370  131  389
645  652  560
496  548  779
910  434  642
793  105  303
232  468  916
932    5  657
782  634  626
429  642  326
946  618  408
760  711  553
561  391  385
614  834  961
585  853  375
188  562  635
775  758  496
300  128  476
747  817  333
288  608  259
410  883  700
142  691  562
222  270  870
654  341  896
548  133  474
 49  712  796
486  607  561
483  920  970
510  553  658
876  682  369
654  744  670
508  888  671
648  111  694
213  954  529
548  879  258
342   15  155
265  880  313
613   36  583
285  774  605
696  776  742
772  230  561
239  304  710
602  387  940
871  107  512
182  321  376
927  392  527
677  124  195
312  270  938
755  308  986
400  779  601
876  843  690
964  719  119
925  665  237
730  719  310
352   86  123
583  801  629
697  340  198
150  635  446
905  183  133
648  654  298
445  743  383
483  628  344
460  822   64
264  872  384
496  291  691
130  742  608
491  590  986
737  317  602
442  179  684
617  256  642
711  688  915
679  804   29
127  869  890
621  677  347
306  486  533
645  198  481
706  855  997
686  743  117
152  947  939
271  251  352
324  621   83
562  745  349
901  797  273
  7   84  696
895  857  751
692  663  805
692  489  122
876  848  930
667  851  155
226  218  502
447  876  635
395   40  430
652  999  312
362  992  135
714  360  668
603  393  858
176   36  470
956  803  884
678  829  391
340  128  810
643  777  545
 71  314  335
705  667  881
119  708  664
480  524  560
432  183  165
983  946  881
788  472  442
386  767  510
864  823  566
764  684  955
155  309  725
459  300  826
627   85  796
497  376  448
827  969  784
408  875  120
764  883  698
 81  590  675
128  549  653
127  606  712
668  989  706
776  440  615
121  840  169
641  648  803
224  671  825
733  419  107
 86  208  359
383  809  426
322  741  122
772   75  577
844  100  782
128  139  344
702  420  230
311  488  724
633  209  661
 33  564  249
459  120  886
493  473  761
252  719  939
506  628  748
673  843  501
124   54  798
421  761  726
521  732   70
395  438  839
600  434  851
464  374   29
598  900  349
817  637  266
558  625  311
503  806  254
527  415  447
131  972  675
816   36  481
870  880  637
215  908  266
973   18  622
973  940  514
463  923  875
472  982  282
868  808  269
544  272  456
961  836   90
130  888  215
974  276  275
309  233  253
973   46  438
842  277  438
366   80  179
419  901  846
 82  907  966
596  354  513
381  362  490
846   11  884
 22  718  970
396  766  862
397   62  598
222  158  646
814  712  225
732  629  623
809  626  692
979  632  811
503  139  372
462  517  811
256  899  609
216  570  483
902  733  385
 89  928    4
887  695  386
 35  568  155
781   58  203
775  604  291
367  692  689
101  158  677
336  580  368
981  337  174
900  880  593
275  613  463
311  907  363
368   83  832
 64  974  980
157  562  421
 12  820  590
160  464  322
245  444  382
  9  312  134
257  306  288
237  449  297
142  600  661
320  363  821
721   84   89
589  509  116
413  594  181
890  477  712
742   65  245
229  432  917
536  189  821
732  401  407
515  210  512
733  778    2
852  451  210
130  360  208
230  408  748
667  499   94
467  112  789
649  764  715
253  908   53
775  878  673
265    5   24
717  434   72
687  428   72
268  436  903
678  450  742
636   40  792
555  104  649
538  608  340
370  525  847
555  830  585
763   92  375
754  898  314
153  560  139
224  663  666
138  344  595
278  448  532
413  492  470
432   98  335
148  795  903
729  903  101
818  186  960
853  631  290
761  170  666
171  582  732
189  731  633
779   20  287
883  726  449
701  139  747
571   29  567
918  166  232
 98  356  853
815  512  449
911  504  671
728  414  257
515  517  657
590  854  517
388  526  831
646  217  989
845  355  289
573  306  156
563   11  456
107  320  601
 37  287  714
167  290  958
198   37  287
896  491  695
712  282  239
223  252  604
524  955  584
883  890  665
818  817  242
518  236  632
410  222  191
310  135  666
983  634  348
671  476  306
986  665  111
109  220  399
717  738  695
764  825  534
616  315  977
628  142  873
 19  287  155
967  255  868
191   80  844
986  220  988
419  521  444
454  916  489
 71  859  500
897  459  731
823  791  216
351  677  556
840  208  612
983  156   22
988  318  633
472  628  495
341  608  343
771  779  528
818  149  422
598   52  436
678  130  285
455  502  177
461  245   81
466  382  258
181  661   64
808  499   22
892  243   76
341  643  531
717  328  856
811  779  683
666  220  797
613  453  417
978  632  462
457  620  387
558  681  351
105  337  432
880   55  818
438   63  136
709  100  700
229  792  280
427  985   53
442  385  325
918  328  642
754  291  642
970   74  973
296   55  952
577  458  924
645  507  523
589  149    6
491  933  297
871  822  303
436  938  577
 98  762  322
368  875  708
607  636  385
488  362  722
642  379  510
271   30  954
338  296  210
125  279  887
614  178  645
268  237  471
578   60  720
776  691  995
814  565  784
 58  358  474
968  573  398
358  613  323
851  694  665
109    4  181
366  741  777
447  747  870
738  460  241
905  694  448
440  901  565
293  278  940
822  276  877
746    2  338
227  915   30
604  733  486
501  359  493
536   79  751
621  623  135
524  547  812
917   11  982
505   55  826
580   55  287
228  805  345
586  101  202
624  829  465
262  645  636
942  775  496
724  942  398
803  499   16
326  565  969
751  977  964
320  725  153
258  772  689
107  421  839
402  399  578
116  927  560
508  685  100
970  581  680
119   98  451
904  580  314
207  186  373
791  286   21
917  199  388
210  549  203
212  270  266
  2  429  355
297  647  659
233  537  895
142  284  332
219  237  361
246  247  401
288   81  328
360  346  279
 21  262  298
343  211   50
637  778  813
820  240   32
660  781  805
638  470  759
779  198  372
158  392  433
  5  274  133
189  346  169
194   74   37
 13  767  447
167  546  364
176  618  336
554  638  712
615  663  776
824   62  142
582  320  499
302  278  545
751  296   71
366   35  493
196  657  381
364  685  134
888  756  128
 17  799  479
872  685  363
879  279  556
665  164   40
264  418  539
627  575  589
978  792  584
662  693    9
988  838  552
870  299   11
141  674  546
460  912  693
216  795  292
531  699  441
207  795  373
719  461  831
571  491  664
142  282   59
 48   89  556
147  278  506
334  990  607
483   42  370
766  978  303
343  336  215
283  745  857
306  587  642
566  764  323
372  267  609
878  505  315
282  877  342
283  369  682
  4  823  926
339  831  891
521   33  942
704  816  318
416  621  503
163  684  625
514  141  646
362   81  368
134  819  425
324  768  190
985  309  356
 41  491  802
997  793  905
976  684  837
368  954  863
878  407   43
216  662  557
 82  425  547
286  486   43
841  595  727
809  169  417
233  566  654
547  419  783
 91  422  981
628    1  945
 83  747  306
399  806  592
346  708  392
813  865  624
516  636   29
592  753  610
440  460  145
457  457  114
 40   19  165
494  659  248
647  950  224
810  965  241
913  630  245
919  652  409
 38  151  355
430  239   96
372  597  360
711  494  370
176  710  108
130  230  503
188  509  421
850  394  702
 68  744  665
919  923  873
`

func TestDay03(t *testing.T) {
	assert := assert.New(t)
	testCases := []tools.TestCase{
		{Details: "Y2016D03 Part 1",
			Input:   "5 10 25",
			Result1: "0"},
		{Details: "Y2016D03 Part 2",
			Input: "101 301 501\n" +
				"102 302 502\n" +
				"103 303 503\n" +
				"201 401 601\n" +
				"202 402 602\n" +
				"203 403 603\n",
			Result2: "6"},
		{Details: "Y2016D03 my input",
			Input:   day03myInput,
			Result1: "869",
			Result2: "1544"},
	}

	for _, tt := range testCases {
		tt.Test(Day03, assert)
	}
}
