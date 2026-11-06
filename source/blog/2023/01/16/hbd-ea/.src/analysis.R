## Libraries
library(ggplot2)

## Read data
setwd("/home/loki/Documents/core/ea/fresh/misc/ea-hbd") ## change to the folder in your computer
data <- read.csv("2020ssc_public.csv", header=TRUE, stringsAsFactors = FALSE)

## Restrict analysis to EAs
data_EAs <- data[data["EAID"] == "Yes",]
View(data_EAs)
n=dim(data_EAs)[1]
n

## Find biodiversity question
colnames(data_EAs)
colnames(data_EAs)[47]

## Process biodiversity question for EAs
tally <- list()
tally$options = c(1:5, "NA")
tally$count = sapply(tally$options, function(x){ sum(data_EAs[47] == x, na.rm = TRUE) })
tally$count[6] = sum(is.na(data_EAs[47]))
tally$count

tally = as.data.frame(tally)
tally

## Plot prevalence of belief within EA

titulo='Prevalence of attitudes towards "human biodiversity"\n amongst EA SlateStarCodex survey respondents in 2020'
subtitulo='"How would you describe your opinion of the the idea of "human biodiversity",\n eg the belief that races differ genetically in socially relevant ways?"\n (1 = very unfavorable, 5 = very favorable), n=993'

(ggplot(data = tally, aes(x =options, y = count)) +
    geom_histogram(
      stat="identity",
      position= position_stack(reverse = TRUE),
      fill="navyblue"
    ))+
  scale_y_continuous(limits = c(0, 300))+
  labs(
    title=titulo, 
    subtitle=subtitulo, 
    x="answers", 
    y="answer count", 
    legend.title = element_blank(),
    legend.text.align = 0
  )+
  theme(
    legend.title = element_blank(), 
    plot.subtitle = element_text(hjust = 0.5), 
    plot.title = element_text(hjust = 0.5),
    legend.position="bottom"
  ) +
  geom_text(aes(label=count,  size = 2), colour="#000000",size=2.5, vjust = -0.5)
  
height=5
width=height*(1+sqrt(5))/2

ggsave("q_hbd_EAs.png" , units="in", width=width, height=height, dpi=800)

## Process biodiversity question for all SSC respondents
tally2 <- list()
tally2$options = c(1:5, "NA")
tally2$count = sapply(tally2$options, function(x){ sum(data[47] == x, na.rm = TRUE) })
tally2$count[6] = sum(is.na(data[47]))
tally2$count
n=dim(data)[1]
n

tally2 = as.data.frame(tally2)
tally2
tally 
## Plot

titulo='Prevalence of attitudes towards "human biodiversity"\n amongst all SlateStarCodex survey respondents in 2020'
subtitulo='"How would you describe your opinion of the the idea of "human biodiversity",\n eg the belief that races differ genetically in socially relevant ways?"\n (1 = very unfavorable, 5 = very favorable), n=7339'

(ggplot(data = tally2, aes(x =options, y = count)) +
    geom_histogram(
      stat="identity",
      position= position_stack(reverse = TRUE),
      fill="navyblue"
    ))+
  scale_y_continuous(limits = c(0, 2000))+
  labs(
    title=titulo, 
    subtitle=subtitulo, 
    x="answers", 
    y="answer count", 
    legend.title = element_blank(),
    legend.text.align = 0
  )+
  theme(
    legend.title = element_blank(), 
    plot.subtitle = element_text(hjust = 0.5), 
    plot.title = element_text(hjust = 0.5),
    legend.position="bottom"
  ) +
  geom_text(aes(label=count,  size = 2), colour="#000000",size=2.5, vjust = -0.5)

height=5
width=height*(1+sqrt(5))/2

ggsave("q_hbd_all.png" , units="in", width=width, height=height, dpi=800)
